package models

type Environment struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Configs     []Config `json:"configs"`
	Secrets     []Secret `json:"secrets"`
}
