package api

import (
	_ "embed"
)

//go:embed schema.openapi.yaml
var configData []byte

// GetConfigSchema returns the embedded open API
// schema for configuration server as a bytes.
func GetConfigSchema() []byte {
	return configData
}
