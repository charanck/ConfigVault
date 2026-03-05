package constants

import (
	"path/filepath"
)

var (
	// ConfigVaultVersion is the current version of the ConfigVault application.
	ConfigVaultVersion = "0.1.0"
	PathToDataFolder   = filepath.Join("../", "data")
)

func SetRoot(rootPath string) {
	PathToDataFolder = filepath.Join(rootPath, "data")
}
