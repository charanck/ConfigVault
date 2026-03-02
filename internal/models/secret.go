package models

type SecretType string

const (
	SecretTypeString  SecretType = "string"
	SecretTypeInt     SecretType = "int"
	SecretTypeBoolean SecretType = "boolean"
	SecretTypeFloat   SecretType = "float"
	SecretTypeObject  SecretType = "object"
	SecretTypeArray   SecretType = "array"
)

type Secret struct {
	Name string     `json:"name"`
	Type SecretType `json:"type"`
}
