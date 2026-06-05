package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Yustinia/kaiten-wall/internal/api"
	"github.com/Yustinia/kaiten-wall/internal/config"
	"github.com/Yustinia/kaiten-wall/internal/download"
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

	result, err := api.FetchWallpapers(&settings)
	if err != nil {
		log.Fatal(err)
	}
	selectedWall := api.SelectRandomWall(result)

	wallLocation, err := download.DownloadWall(selectedWall, settings.General.DefaultPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Wallpaper Location: %s", wallLocation)
}
