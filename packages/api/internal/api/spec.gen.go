// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
