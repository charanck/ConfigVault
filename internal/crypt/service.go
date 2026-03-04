package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/charanck/ConfigVault/config"
)

type CryptService struct {
	config *config.Config
}

func NewCryptService(cfg *config.Config) *CryptService {
	return &CryptService{
		config: cfg,
	}
}

func (s *CryptService) GenerateKeyPair() (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", err
	}

	privBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privBytes,
	})

	pubBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}
	pubPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	})

	return string(pubPEM), string(privPEM), nil
}

func decodePublicKey(publicKeyStr string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pub, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}
	return pub, nil
}

func (s *CryptService) Encrypt(value string, publicKey string) (string, error) {
	decodedPublicKey, err := decodePublicKey(publicKey)
	if err != nil {
		return "", err
	}
	label := []byte("ConfigVault")
	encryptedData, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, decodedPublicKey, []byte(value), label)
	if err != nil {
		return "", err
	}
	return string(encryptedData), nil
}
