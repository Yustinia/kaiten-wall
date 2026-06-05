package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Yustinia/kaiten-wall/internal/config"
	"github.com/Yustinia/kaiten-wall/internal/fs"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("home not found: %v", err)
	}

	configPath := filepath.Join(homeDir, ".config", "kaiten-wall", "config.toml")
	if err = fs.IsFileExist(configPath); err != nil {
		log.Fatalf("config not found: %v", err)
	}

	settings, err := config.ParseSettings(configPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", settings)
}
