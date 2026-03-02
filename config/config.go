package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type GitProviderType string

const (
	GitProviderTypeGitHub    GitProviderType = "github"
	GitProviderTypeGitLab    GitProviderType = "gitlab"
	GitProviderTypeBitbucket GitProviderType = "bitbucket"
)

type EnvType string

const (
	EnvTypeDevelopment EnvType = "development"
	EnvTypeStaging     EnvType = "staging"
	EnvTypeProduction  EnvType = "production"
)

type Config struct {
	GitProviderType GitProviderType `envconfig:"GIT_PROVIDER_TYPE" default:"github"`

	GitHubToken          string `envconfig:"GITHUB_TOKEN"`
	GitLabToken          string `envconfig:"GITLAB_TOKEN"`
	BitbucketUsername    string `envconfig:"BITBUCKET_USERNAME"`
	BitbucketAppPassword string `envconfig:"BITBUCKET_APP_PASSWORD"`

	GitRepoName string  `envconfig:"GIT_REPO_NAME" required:"true"`
	RedisURL    string  `envconfig:"REDIS_URL" default:"redis://localhost:6379"`
	Env         EnvType `envconfig:"ENV" default:"development"`
	Port        string  `envconfig:"PORT" default:"8080"`

	SOPSAgeKey string `envconfig:"SOPS_AGE_KEY"`
}

func LoadConfig() (*Config, error) {
	var cfg Config

	env := os.Getenv("ENV")
	if env == "" || env == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found, proceeding with system env vars")
		}
	}

	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return &cfg, nil
}
