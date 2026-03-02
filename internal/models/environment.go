package models

type Environment struct {
	Id            string         `json:"id"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	ConfigValues  map[string]any `json:"configs"`
	SecretsValues map[string]any `json:"secrets"` // Secret values will be stored in encrypted format
}
