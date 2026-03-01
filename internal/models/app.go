package models

type App struct {
	Id           string        `json:"id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Tags         []string      `json:"tags"`
	ClientID     string        `json:"client_id"`
	ClientSecret string        `json:"client_secret"`
	Environments []Environment `json:"environments"`
}
