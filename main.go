package main

import (
	"fmt"
	"os"

	"github.com/charanck/ConfigVault/common/constants"
	"github.com/charanck/ConfigVault/config"
	"github.com/charanck/ConfigVault/internal/app"
	"github.com/charanck/ConfigVault/internal/crypt"
)

func main() {
	pwd, _ := os.Getwd()
	constants.SetRoot(pwd)

	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	cryptService := crypt.NewCryptService(config)
	appService := app.NewAppService(config, cryptService)

	createdApp, clientSecret, privateCertificate, err := appService.CreateApp(app.CreateAppRequest{
		Name:        "My First App",
		Description: "This is a test app",
		Tags:        []string{"test", "example"},
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("App created successfully")
	fmt.Println("ClientSecret: ", clientSecret)
	fmt.Println("PrivateCertificate: ", privateCertificate)
	fmt.Println("App Details: ", createdApp)
}
