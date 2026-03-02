package main

import (
	"fmt"

	"github.com/charanck/ConfigVault/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	fmt.Println("Config loaded successfully:", config)
}
