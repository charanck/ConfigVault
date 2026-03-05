package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	"github.com/alexedwards/argon2id"
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

func (s *CryptService) decodePublicKey(publicKeyStr string) (*rsa.PublicKey, error) {
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
	decodedPublicKey, err := s.decodePublicKey(publicKey)
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

func (s *CryptService) Decrypt(encryptedValue string, privateKeyStr string) (string, error) {
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		return "", fmt.Errorf("failed to parse PEM block")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	label := []byte("ConfigVault")
	decryptedData, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, []byte(encryptedValue), label)
	if err != nil {
		return "", err
	}
	return string(decryptedData), nil
}

func (s *CryptService) GenerateClientID() (string, error) {
	clientIDBytes := make([]byte, 16)
	_, err := rand.Read(clientIDBytes)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", clientIDBytes), nil
}

func (s *CryptService) GenerateClientSecret() (string, error) {
	b := make([]byte, 32)

	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("entropy source failure: %w", err)
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

func (s *CryptService) HashSecret(plainSecret string) (string, error) {
	params := &argon2id.Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	hash, err := argon2id.CreateHash(plainSecret, params)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (s *CryptService) VerifySecret(plainSecret, hashedSecret string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(plainSecret, hashedSecret)
	if err != nil {
		return false, err
	}
	return match, nil
}
