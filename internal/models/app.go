package models

type App struct {
	Id                string            `json:"id"`
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	Tags              []string          `json:"tags"`
	ClientID          string            `json:"client_id"`
	ClientSecret      string            `json:"client_secret"`
	PublicCertificate string            `json:"public_certificate"`
	Configs           map[string]Config `json:"configs"`
	Secrets           map[string]Secret `json:"secrets"`
	Environments      []Environment     `json:"environments"`
}
