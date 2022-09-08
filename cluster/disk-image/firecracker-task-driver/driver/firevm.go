/* Firecracker-task-driver is a task driver for Hashicorp's nomad that allows
 * to create microvms using AWS Firecracker vmm
 * Copyright (C) 2019  Carlos Neira cneirabustos@gmail.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License")
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 */

package firevm

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/cneira/firecracker-task-driver/driver/client/client"
	"github.com/cneira/firecracker-task-driver/driver/client/client/operations"
	"github.com/cneira/firecracker-task-driver/driver/client/models"
	firecracker "github.com/firecracker-microvm/firecracker-go-sdk"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/nomad/plugins/drivers"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const (
	// containerMonitorIntv is the interval at which the driver checks if the
	// firecracker micro-vm is still running
	containerMonitorIntv = 4 * time.Second

	editIDName          = "edit_id"
	buildIDName         = "build_id"
	templateIDName      = "template_id"
	templateBuildIDName = "template_build_id"
	rootfsName          = "rootfs.ext4"
	snapfileName        = "snapfile"
	memfileName         = "memfile"

	editDirName  = "edit"
	buildDirName = "builds"
)

type vminfo struct {
	Machine *firecracker.Machine
	Info    Instance_info
}
type Instance_info struct {
	AllocId              string
	Pid                  string
	SnapshotRootPath     string
	EditID               *string
	SocketPath           string
	CodeSnippetID        string
	CodeSnippetDirectory string
	BuildDirPath         string
	Cmd                  *exec.Cmd
}

func newFirecrackerClient(socketPath string) *client.Firecracker {
	httpClient := client.NewHTTPClient(strfmt.NewFormats())

	transport := firecracker.NewUnixSocketTransport(socketPath, nil, false)
	httpClient.SetTransport(transport)

	return httpClient
}

func loadSnapshot(ctx context.Context, socketPath, snapshotRootPath string, d *Driver) error {
	childCtx, childSpan := d.tracer.Start(ctx, "load-snapshot", trace.WithAttributes(
		attribute.String("socket_path", socketPath),
		attribute.String("snapshot_root_path", snapshotRootPath),
	))
	defer childSpan.End()

	httpClient := newFirecrackerClient(socketPath)
	d.ReportEvent(childCtx, "created FC socket client")

	memfilePath := filepath.Join(snapshotRootPath, memfileName)
	snapfilePath := filepath.Join(snapshotRootPath, snapfileName)

	childSpan.SetAttributes(
		attribute.String("memfile_path", memfilePath),
		attribute.String("snapfile_path", snapfilePath),
	)

	backendType := models.MemoryBackendBackendTypeFile
	snapshotConfig := operations.LoadSnapshotParams{
		Context: ctx,
		Body: &models.SnapshotLoadParams{
			ResumeVM:            true,
			EnableDiffSnapshots: false,
			MemBackend: &models.MemoryBackend{
				BackendPath: &memfilePath,
				BackendType: &backendType,
			},
			SnapshotPath: &snapfilePath,
		},
	}

	_, err := httpClient.Operations.LoadSnapshot(&snapshotConfig)
	if err != nil {
		childSpan.RecordError(err)
		childSpan.SetStatus(codes.Error, "critical error")
		return err
	}
	d.ReportEvent(childCtx, "snapshot loaded")

	return nil
}

func (d *Driver) initializeFC(
	ctx context.Context,
	cfg *drivers.TaskConfig,
	taskConfig TaskConfig,
	slot *IPSlot,
	fcEnvsDisk string,
	editEnabled bool,
) (*vminfo, error) {
	childCtx, childSpan := d.tracer.Start(ctx, "initialize-fc", trace.WithAttributes(
		attribute.String("session_id", slot.SessionID),
		attribute.Int("ip_slot_index", slot.SlotIdx),
	))
	defer childSpan.End()

	opts := newOptions()

	opts.FcLogFifo = path.Join(cfg.AllocDir, "fc-log-fifo")
	syscall.Mkfifo(opts.FcLogFifo, 0644)
	opts.FcMetricsFifo = path.Join(cfg.AllocDir, "fc-metrics-fifo")
	syscall.Mkfifo(opts.FcMetricsFifo, 0644)
	fcCfg, err := opts.getFirecrackerConfig(cfg.AllocID)
	if err != nil {
		errMsg := fmt.Errorf("error assembling FC config: %v", err)
		d.ReportCriticalError(childCtx, errMsg)
		return nil, errMsg
	}

	vmmChildCtx, vmmChildSpan := d.tracer.Start(
		childCtx,
		"fc-vmm",
		trace.WithAttributes(attribute.String("fc_log_fifo", opts.FcLogFifo)),
		trace.WithAttributes(attribute.String("fc_metrics_fifo", opts.FcMetricsFifo)),
	)
	defer vmmChildSpan.End()

	vmmCtx, vmmCancel := context.WithCancel(vmmChildCtx)
	defer vmmCancel()

	d.logger.Info("Starting firecracker", "driver_initialize_container", hclog.Fmt("%v+", opts))
	logger := log.New()
	log.SetLevel(log.DebugLevel)
	logger.SetLevel(log.DebugLevel)

	otelHook := NewOtelHook(vmmChildSpan)
	logger.AddHook(otelHook)

	machineOpts := []firecracker.Opt{
		firecracker.WithLogger(log.NewEntry(logger)),
	}

	os.MkdirAll(slot.SessionTmpOverlay(), 0777)
	os.MkdirAll(slot.SessionTmpWorkdir(), 0777)

	codeSnippetEnvPath := filepath.Join(fcEnvsDisk, taskConfig.CodeSnippetID)

	var buildDirPath string
	var snapshotRootPath string
	var mountCmd string
	var editID string

	if editEnabled {
		// Use the shared edit sessions
		codeSnippetEditPath := filepath.Join(codeSnippetEnvPath, editDirName)

		buildIDSrc := filepath.Join(codeSnippetEnvPath, buildIDName)
		buildIDDest := filepath.Join(codeSnippetEditPath, buildIDName)

		templateIDSrc := filepath.Join(codeSnippetEnvPath, templateIDName)
		templateIDDest := filepath.Join(codeSnippetEditPath, templateIDName)

		templateBuildIDSrc := filepath.Join(codeSnippetEnvPath, templateBuildIDName)
		templateBuildIDDest := filepath.Join(codeSnippetEditPath, templateBuildIDName)

		editIDPath := filepath.Join(codeSnippetEditPath, editIDName)

		os.MkdirAll(codeSnippetEditPath, 0777)

		if _, err := os.Stat(editIDPath); err == nil {
			// If the edit_file exists we expect that the other files will exists too (we are creating te edit last)
			data, err := os.ReadFile(editIDPath)
			if err != nil {
				return nil, fmt.Errorf("failed reading edit id for the code snippet %s: %v", taskConfig.CodeSnippetID, err)
			}
			editID = string(data)

			snapshotRootPath = filepath.Join(codeSnippetEditPath, editID)
		} else {
			// Link the fc files from the root CS directory and create edit_id
			editID = uuid.New().String()

			snapshotRootPath = filepath.Join(codeSnippetEditPath, editID)
			os.MkdirAll(snapshotRootPath, 0777)

			rootfsSrc := filepath.Join(codeSnippetEnvPath, rootfsName)
			rootfsDest := filepath.Join(codeSnippetEditPath, editID, rootfsName)

			snapfileSrc := filepath.Join(codeSnippetEnvPath, snapfileName)
			snapfileDest := filepath.Join(codeSnippetEditPath, editID, snapfileName)

			memfileSrc := filepath.Join(codeSnippetEnvPath, memfileName)
			memfileDest := filepath.Join(codeSnippetEditPath, editID, memfileName)

			os.Link(rootfsSrc, rootfsDest)
			os.Link(snapfileSrc, snapfileDest)
			os.Link(memfileSrc, memfileDest)
			os.Link(buildIDSrc, buildIDDest)
			os.Link(templateIDSrc, templateIDDest)
			os.Link(templateBuildIDSrc, templateBuildIDDest)

			err := os.WriteFile(editIDPath, []byte(editID), 0777)
			if err != nil {
				fmt.Printf("Unable to create edit_id file: %v", err)
			}
		}

		if _, err := os.Stat(buildIDDest); err != nil {
			// If the build_id file does not exist this envs is templated - we check to which env the template points to and use that as our new "virtual" env
			data, err := os.ReadFile(templateIDDest)
			if err != nil {
				return nil, fmt.Errorf("failed reading template id for the code snippet %s: %v", taskConfig.CodeSnippetID, err)
			}
			templateID := string(data)

			templateEnvPath := filepath.Join(fcEnvsDisk, templateID)

			data, err = os.ReadFile(templateBuildIDDest)
			if err != nil {
				return nil, fmt.Errorf("failed reading build id for the template %s of code snippet %s: %v", templateID, taskConfig.CodeSnippetID, err)
			}
			templateBuildID := string(data)
			buildDirPath = filepath.Join(templateEnvPath, buildDirName, templateBuildID)
		} else {
			// build_id is present so this is a non-templated session
			data, err := os.ReadFile(buildIDDest)
			if err != nil {
				return nil, fmt.Errorf("failed reading build id for the code snippet %s: %v", taskConfig.CodeSnippetID, err)
			}
			buildID := string(data)
			buildDirPath = filepath.Join(codeSnippetEnvPath, buildDirName, buildID)
		}

		os.MkdirAll(buildDirPath, 0777)

		mountCmd = fmt.Sprintf(
			"mount -t overlay overlay -o lowerdir=%s,upperdir=%s,workdir=%s %s && ",
			snapshotRootPath,
			slot.SessionTmpOverlay(),
			slot.SessionTmpWorkdir(),
			buildDirPath,
		)
	} else {
		// Mount overlay
		snapshotRootPath = codeSnippetEnvPath

		templateIDPath := filepath.Join(codeSnippetEnvPath, templateIDName)
		templateBuildIDPath := filepath.Join(codeSnippetEnvPath, templateBuildIDName)
		buildIDPath := filepath.Join(codeSnippetEnvPath, buildIDName)

		if _, err := os.Stat(buildIDPath); err != nil {
			// If the build_id file does not exist this envs is templated - we check to which env the template points to and use that as our new "virtual" env
			data, err := os.ReadFile(templateIDPath)
			if err != nil {
				return nil, fmt.Errorf("failed reading template id for the code snippet %s: %v", taskConfig.CodeSnippetID, err)
			}
			templateID := string(data)

			templateEnvPath := filepath.Join(fcEnvsDisk, templateID)

			data, err = os.ReadFile(templateBuildIDPath)
			if err != nil {
				return nil, fmt.Errorf("failed reading build id for the template %s of code snippet %s: %v", templateID, taskConfig.CodeSnippetID, err)
			}
			templateBuildID := string(data)
			buildDirPath = filepath.Join(templateEnvPath, buildDirName, templateBuildID)
		} else {
			// build_id is present and this is a normal non-templated and non-edit session
			data, err := os.ReadFile(buildIDPath)
			if err != nil {
				return nil, fmt.Errorf("failed reading build id for the code snippet %s: %v", taskConfig.CodeSnippetID, err)
			}
			buildID := string(data)
			buildDirPath = filepath.Join(codeSnippetEnvPath, buildDirName, buildID)
		}

		os.MkdirAll(buildDirPath, 0777)

		mountCmd = fmt.Sprintf(
			"mount -t overlay overlay -o lowerdir=%s,upperdir=%s,workdir=%s %s && ",
			snapshotRootPath,
			slot.SessionTmpOverlay(),
			slot.SessionTmpWorkdir(),
			buildDirPath,
		)
	}

	fcCmd := fmt.Sprintf("/usr/bin/firecracker --api-sock %s", fcCfg.SocketPath)
	inNetNSCmd := fmt.Sprintf("ip netns exec %s ", slot.NamespaceID())

	cmd := exec.CommandContext(childCtx, "unshare", "-pfm", "--kill-child", "--", "bash", "-c", mountCmd+inNetNSCmd+fcCmd)
	cmd.Stderr = nil

	machineOpts = append(machineOpts, firecracker.WithProcessRunner(cmd))

	vmmLogsReader, vmmLogsWriter := io.Pipe()

	go func() {
		scanner := bufio.NewScanner(vmmLogsReader)

		for scanner.Scan() {
			line := scanner.Text()

			vmmChildSpan.AddEvent("log message",
				trace.WithAttributes(
					attribute.String("message", line),
				),
			)
		}

		vmmLogsReader.Close()
	}()

	prebootFcConfig := firecracker.Config{
		DisableValidation: true,
		MmdsAddress:       fcCfg.MmdsAddress,
		Seccomp:           fcCfg.Seccomp,
		ForwardSignals:    fcCfg.ForwardSignals,
		VMID:              fcCfg.VMID,
		MachineCfg:        fcCfg.MachineCfg,
		VsockDevices:      fcCfg.VsockDevices,
		FifoLogWriter:     vmmLogsWriter,
		Drives:            fcCfg.Drives,
		KernelArgs:        fcCfg.KernelArgs,
		InitrdPath:        fcCfg.InitrdPath,
		KernelImagePath:   fcCfg.KernelImagePath,
		MetricsFifo:       fcCfg.MetricsFifo,
		MetricsPath:       fcCfg.MetricsPath,
		LogLevel:          fcCfg.LogLevel,
		LogFifo:           fcCfg.LogFifo,
		LogPath:           fcCfg.LogPath,
		SocketPath:        fcCfg.SocketPath,
	}

	m, err := firecracker.NewMachine(vmmCtx, prebootFcConfig, machineOpts...)
	if err != nil {
		errMsg := fmt.Errorf("failed creating machine: %v", err)

		childSpan.RecordError(errMsg)
		childSpan.SetStatus(codes.Error, "critical error")
		return nil, errMsg
	}
	d.ReportEvent(childCtx, "created vmm")

	m.Handlers.Validation = m.Handlers.Validation.Clear()
	m.Handlers.FcInit =
		m.Handlers.FcInit.Clear().
			Append(
				firecracker.StartVMMHandler,
				firecracker.BootstrapLoggingHandler,
			)

	err = m.Handlers.Run(childCtx, m)
	if err != nil {
		errMsg := fmt.Errorf("failed to start preboot FC: %v", err)

		childSpan.RecordError(errMsg)
		childSpan.SetStatus(codes.Error, "critical error")
		return nil, errMsg
	}
	d.ReportEvent(childCtx, "started FC in preboot")

	if err := loadSnapshot(childCtx, fcCfg.SocketPath, snapshotRootPath, d); err != nil {
		m.StopVMM()
		errMsg := fmt.Errorf("failed to load snapshot: %v", err)

		childSpan.RecordError(errMsg)
		childSpan.SetStatus(codes.Error, "critical error")
		return nil, errMsg
	}
	d.ReportEvent(childCtx, "loaded snapshot")

	defer func() {
		if err != nil {
			stopErr := m.StopVMM()
			if stopErr != nil {
				logger.Error(fmt.Sprintf("Failed stopping machine after error: %+v", stopErr))
			}
		}
	}()

	if opts.validMetadata != nil {
		m.SetMetadata(vmmCtx, opts.validMetadata)
	}

	pid, errpid := m.PID()
	if errpid != nil {
		errMsg := fmt.Errorf("failed getting pid for machine: %v", errpid)

		childSpan.RecordError(errMsg)
		childSpan.SetStatus(codes.Error, "critical error")
		return nil, errMsg
	}

	info := Instance_info{
		Cmd:                  cmd,
		AllocId:              cfg.AllocID,
		Pid:                  strconv.Itoa(pid),
		SnapshotRootPath:     snapshotRootPath,
		EditID:               &editID,
		SocketPath:           fcCfg.SocketPath,
		CodeSnippetID:        taskConfig.CodeSnippetID,
		CodeSnippetDirectory: codeSnippetEnvPath,
		BuildDirPath:         buildDirPath,
	}

	return &vminfo{Machine: m, Info: info}, nil
}
