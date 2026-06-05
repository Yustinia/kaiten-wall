package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Yustinia/kaiten-wall/internal/api"
	"github.com/Yustinia/kaiten-wall/internal/config"
	"github.com/Yustinia/kaiten-wall/internal/daemon"
	"github.com/Yustinia/kaiten-wall/internal/defaults"
	"github.com/Yustinia/kaiten-wall/internal/download"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("home not found: %v", err)
	}

	configDir := filepath.Join(homeDir, ".config", "kaiten-wall")
	configPath := filepath.Join(homeDir, ".config", "kaiten-wall", "config.toml")

	if err = defaults.FirstLaunch(configDir, configPath); err != nil {
		log.Fatal(err)
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

	err = daemon.RunAwww(settings.General.UseDaemon, wallLocation, &settings.Awww)
	if err != nil {
		log.Fatal(err)
	}
}
