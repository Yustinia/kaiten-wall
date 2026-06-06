package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/Yustinia/kaiten-wall/internal/api"
	"github.com/Yustinia/kaiten-wall/internal/config"
	"github.com/Yustinia/kaiten-wall/internal/daemon"
	"github.com/Yustinia/kaiten-wall/internal/defaults"
	"github.com/Yustinia/kaiten-wall/internal/download"
	"github.com/Yustinia/kaiten-wall/internal/theming"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("home not found: %v", err)
	}

	configDir := filepath.Join(homeDir, ".config", "kaiten-wall")
	configPath := filepath.Join(homeDir, ".config", "kaiten-wall", "config.toml")

	if err = defaults.FirstLaunch(configDir, configPath); err != nil {
		if errors.Is(err, defaults.ErrFirstLaunch) {
			log.Println(err)
			os.Exit(0)
		}
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
	selectedWall, err := api.SelectRandomWall(result)
	if err != nil {
		log.Fatal(err)
	}

	wallLocation, err := download.DownloadWall(selectedWall, settings.General.DefaultPath)
	if err != nil {
		log.Fatal(err)
	}

	switch settings.General.UseDaemon {
	case "awww":
		err = daemon.RunAwww(wallLocation, &settings.Awww)
	default:
		log.Fatalf("unknown daemon: %q", settings.General.UseDaemon)
	}
	if err != nil {
		log.Fatal(err)
	}

	if settings.General.UseThemer != "" {
		switch settings.General.UseThemer {
		case "matugen":
			err = theming.ApplyMatugen(wallLocation, &settings.Matugen)
		default:
			log.Fatalf("unknown themer: %q", settings.General.UseThemer)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
