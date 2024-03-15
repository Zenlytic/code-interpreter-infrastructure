// Code generated by ent, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/env"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/envbuild"
	"github.com/google/uuid"
)

// EnvBuild is the model entity for the EnvBuild schema.
type EnvBuild struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// FinishedAt holds the value of the "finished_at" field.
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	// EnvID holds the value of the "env_id" field.
	EnvID *string `json:"env_id,omitempty"`
	// Status holds the value of the "status" field.
	Status envbuild.Status `json:"status,omitempty"`
	// StartCmd holds the value of the "start_cmd" field.
	StartCmd *string `json:"start_cmd,omitempty"`
	// Vcpu holds the value of the "vcpu" field.
	Vcpu int64 `json:"vcpu,omitempty"`
	// RAMMB holds the value of the "ram_mb" field.
	RAMMB int64 `json:"ram_mb,omitempty"`
	// FreeDiskSizeMB holds the value of the "free_disk_size_mb" field.
	FreeDiskSizeMB int64 `json:"free_disk_size_mb,omitempty"`
	// TotalDiskSizeMB holds the value of the "total_disk_size_mb" field.
	TotalDiskSizeMB *int64 `json:"total_disk_size_mb,omitempty"`
	// KernelVersion holds the value of the "kernel_version" field.
	KernelVersion string `json:"kernel_version,omitempty"`
	// FirecrackerVersion holds the value of the "firecracker_version" field.
	FirecrackerVersion string `json:"firecracker_version,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EnvBuildQuery when eager-loading is set.
	Edges        EnvBuildEdges `json:"edges"`
	selectValues sql.SelectValues
}

// EnvBuildEdges holds the relations/edges for other nodes in the graph.
type EnvBuildEdges struct {
	// Env holds the value of the env edge.
	Env *Env `json:"env,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EnvOrErr returns the Env value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EnvBuildEdges) EnvOrErr() (*Env, error) {
	if e.loadedTypes[0] {
		if e.Env == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: env.Label}
		}
		return e.Env, nil
	}
	return nil, &NotLoadedError{edge: "env"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EnvBuild) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case envbuild.FieldVcpu, envbuild.FieldRAMMB, envbuild.FieldFreeDiskSizeMB, envbuild.FieldTotalDiskSizeMB:
			values[i] = new(sql.NullInt64)
		case envbuild.FieldEnvID, envbuild.FieldStatus, envbuild.FieldStartCmd, envbuild.FieldKernelVersion, envbuild.FieldFirecrackerVersion:
			values[i] = new(sql.NullString)
		case envbuild.FieldCreatedAt, envbuild.FieldUpdatedAt, envbuild.FieldFinishedAt:
			values[i] = new(sql.NullTime)
		case envbuild.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EnvBuild fields.
func (eb *EnvBuild) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case envbuild.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				eb.ID = *value
			}
		case envbuild.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				eb.CreatedAt = value.Time
			}
		case envbuild.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				eb.UpdatedAt = value.Time
			}
		case envbuild.FieldFinishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field finished_at", values[i])
			} else if value.Valid {
				eb.FinishedAt = new(time.Time)
				*eb.FinishedAt = value.Time
			}
		case envbuild.FieldEnvID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field env_id", values[i])
			} else if value.Valid {
				eb.EnvID = new(string)
				*eb.EnvID = value.String
			}
		case envbuild.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				eb.Status = envbuild.Status(value.String)
			}
		case envbuild.FieldStartCmd:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field start_cmd", values[i])
			} else if value.Valid {
				eb.StartCmd = new(string)
				*eb.StartCmd = value.String
			}
		case envbuild.FieldVcpu:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field vcpu", values[i])
			} else if value.Valid {
				eb.Vcpu = value.Int64
			}
		case envbuild.FieldRAMMB:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ram_mb", values[i])
			} else if value.Valid {
				eb.RAMMB = value.Int64
			}
		case envbuild.FieldFreeDiskSizeMB:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field free_disk_size_mb", values[i])
			} else if value.Valid {
				eb.FreeDiskSizeMB = value.Int64
			}
		case envbuild.FieldTotalDiskSizeMB:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_disk_size_mb", values[i])
			} else if value.Valid {
				eb.TotalDiskSizeMB = new(int64)
				*eb.TotalDiskSizeMB = value.Int64
			}
		case envbuild.FieldKernelVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kernel_version", values[i])
			} else if value.Valid {
				eb.KernelVersion = value.String
			}
		case envbuild.FieldFirecrackerVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field firecracker_version", values[i])
			} else if value.Valid {
				eb.FirecrackerVersion = value.String
			}
		default:
			eb.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EnvBuild.
// This includes values selected through modifiers, order, etc.
func (eb *EnvBuild) Value(name string) (ent.Value, error) {
	return eb.selectValues.Get(name)
}

// QueryEnv queries the "env" edge of the EnvBuild entity.
func (eb *EnvBuild) QueryEnv() *EnvQuery {
	return NewEnvBuildClient(eb.config).QueryEnv(eb)
}

// Update returns a builder for updating this EnvBuild.
// Note that you need to call EnvBuild.Unwrap() before calling this method if this EnvBuild
// was returned from a transaction, and the transaction was committed or rolled back.
func (eb *EnvBuild) Update() *EnvBuildUpdateOne {
	return NewEnvBuildClient(eb.config).UpdateOne(eb)
}

// Unwrap unwraps the EnvBuild entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (eb *EnvBuild) Unwrap() *EnvBuild {
	_tx, ok := eb.config.driver.(*txDriver)
	if !ok {
		panic("models: EnvBuild is not a transactional entity")
	}
	eb.config.driver = _tx.drv
	return eb
}

// String implements the fmt.Stringer.
func (eb *EnvBuild) String() string {
	var builder strings.Builder
	builder.WriteString("EnvBuild(")
	builder.WriteString(fmt.Sprintf("id=%v, ", eb.ID))
	builder.WriteString("created_at=")
	builder.WriteString(eb.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(eb.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := eb.FinishedAt; v != nil {
		builder.WriteString("finished_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := eb.EnvID; v != nil {
		builder.WriteString("env_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", eb.Status))
	builder.WriteString(", ")
	if v := eb.StartCmd; v != nil {
		builder.WriteString("start_cmd=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("vcpu=")
	builder.WriteString(fmt.Sprintf("%v", eb.Vcpu))
	builder.WriteString(", ")
	builder.WriteString("ram_mb=")
	builder.WriteString(fmt.Sprintf("%v", eb.RAMMB))
	builder.WriteString(", ")
	builder.WriteString("free_disk_size_mb=")
	builder.WriteString(fmt.Sprintf("%v", eb.FreeDiskSizeMB))
	builder.WriteString(", ")
	if v := eb.TotalDiskSizeMB; v != nil {
		builder.WriteString("total_disk_size_mb=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("kernel_version=")
	builder.WriteString(eb.KernelVersion)
	builder.WriteString(", ")
	builder.WriteString("firecracker_version=")
	builder.WriteString(eb.FirecrackerVersion)
	builder.WriteByte(')')
	return builder.String()
}

// EnvBuilds is a parsable slice of EnvBuild.
type EnvBuilds []*EnvBuild
