package sops

import "github.com/charanck/ConfigVault/config"

// SOPSService is responsible for handling all operations related to SOPS encryption and decryption.
type SOPSService struct {
	config *config.Config
}

func NewSopsService(cfg *config.Config) *SOPSService {
	return &SOPSService{
		config: cfg,
	}
}
