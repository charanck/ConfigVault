package models

type ConfigType string

const (
	ConfigTypeString  ConfigType = "string"
	ConfigTypeInt     ConfigType = "int"
	ConfigTypeBoolean ConfigType = "boolean"
	ConfigTypeFloat   ConfigType = "float"
	ConfigTypeObject  ConfigType = "object"
	ConfigTypeArray   ConfigType = "array"
)

type Config struct {
	Name string     `json:"name"`
	Type ConfigType `json:"type"`
}
