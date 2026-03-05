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

func (e *EnvType) Set(value string) error {
	switch EnvType(value) {
	case EnvTypeDevelopment, EnvTypeStaging, EnvTypeProduction:
		*e = EnvType(value)
		return nil
	default:
		return fmt.Errorf("invalid ENV value: %s", value)
	}
}

type Config struct {
	GitProviderType GitProviderType `envconfig:"GIT_PROVIDER_TYPE" default:"github"`

	GitHubToken          string `envconfig:"GITHUB_TOKEN"`
	GitLabToken          string `envconfig:"GITLAB_TOKEN"`
	BitbucketUsername    string `envconfig:"BITBUCKET_USERNAME"`
	BitbucketAppPassword string `envconfig:"BITBUCKET_APP_PASSWORD"`

	GitRepoName string  `envconfig:"GIT_REPO_NAME" required:"true"`
	RedisURL    string  `envconfig:"REDIS_URL" default:"redis://localhost:6379"`
	Env         EnvType `envconfig:"ENV" default:"development" required:"true"`
	Port        string  `envconfig:"PORT" default:"8080" required:"true"`

	SOPSAgePublicKey  string `envconfig:"SOPS_AGE_PUBLIC_KEY" required:"true"`
	SOPSAgePrivateKey string `envconfig:"SOPS_AGE_PRIVATE_KEY"`
}

var sensitiveKeys = []string{
	"SOPS_AGE_PRIVATE_KEY",
	"GITHUB_TOKEN",
	"GITLAB_TOKEN",
	"BITBUCKET_APP_PASSWORD",
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

	// Wipe confidential env vars
	for _, key := range sensitiveKeys {
		os.Unsetenv(key)
	}

	return &cfg, nil
}
