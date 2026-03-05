package app

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/charanck/ConfigVault/common/constants"
	"github.com/charanck/ConfigVault/config"
	"github.com/charanck/ConfigVault/internal/crypt"
	"github.com/charanck/ConfigVault/internal/models"
	"github.com/google/uuid"
)

type AppService struct {
	config       *config.Config
	cryptService *crypt.CryptService
}

func NewAppService(cfg *config.Config, cryptService *crypt.CryptService) *AppService {
	return &AppService{
		config:       cfg,
		cryptService: cryptService,
	}
}

type CreateAppRequest struct {
	Name        string
	Description string
	Tags        []string
}

func (s *AppService) CreateApp(data CreateAppRequest) (models.App, string, string, error) {
	publicCertificate, privateCertificate, err := s.cryptService.GenerateKeyPair()
	if err != nil {
		return models.App{}, "", "", err
	}
	clientID, err := s.cryptService.GenerateClientID()
	if err != nil {
		return models.App{}, "", "", err
	}
	clientSecret, err := s.cryptService.GenerateClientSecret()
	if err != nil {
		return models.App{}, "", "", err
	}
	hashedClientSecret, err := s.cryptService.HashSecret(clientSecret)
	if err != nil {
		return models.App{}, "", "", err
	}

	newApp := models.App{
		Id:                uuid.NewString(),
		Name:              data.Name,
		Description:       data.Description,
		Tags:              data.Tags,
		PublicCertificate: publicCertificate,
		ClientID:          clientID,
		ClientSecret:      hashedClientSecret,
		Configs:           make(map[string]models.Config),
		Secrets:           make(map[string]models.Secret),
		Environments:      []models.Environment{},
	}

	file, err := os.Create(fmt.Sprintf("%s/apps/%s.json", constants.PathToDataFolder, newApp.Id))
	if err != nil {
		return models.App{}, "", "", err
	}
	defer file.Close()
	jsonBytes, err := json.Marshal(newApp)
	if err != nil {
		return models.App{}, "", "", err
	}
	_, err = file.Write(jsonBytes)
	if err != nil {
		return models.App{}, "", "", err
	}

	// TODO: commit the changes to git
	return newApp, clientSecret, privateCertificate, nil
}
