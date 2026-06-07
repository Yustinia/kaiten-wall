package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Yustinia/kaiten-wall/internal/api"
	"github.com/Yustinia/kaiten-wall/internal/config"
	"github.com/Yustinia/kaiten-wall/internal/daemon"
	"github.com/Yustinia/kaiten-wall/internal/download"
	"github.com/Yustinia/kaiten-wall/internal/fs"
	"github.com/Yustinia/kaiten-wall/internal/theming"
)

func TestMain(t *testing.T) {
	projDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("project root not found: %v", err)
	}

	configPath := filepath.Join(projDir, "..", "internal", "defaults", "config.example.toml")
	if err = fs.IsFileExist(configPath); err != nil {
		t.Fatalf("config not found: %v", err)
	}

	settings, err := config.ParseSettings(configPath)
	if err != nil {
		t.Fatal(err)
	}

	result, err := api.FetchWallpapers(&settings)
	if err != nil {
		t.Fatal(err)
	}

	selectedWall, err := api.SelectRandomWall(result)
	if err != nil {
		t.Fatal(err)
	}

	wallLocation, err := download.DownloadWall(selectedWall, settings.General.DefaultPath)
	if err != nil {
		t.Fatal(err)
	}

	switch settings.General.UseDaemon {
	case "awww":
		err = daemon.RunAwww(wallLocation, &settings.Awww)
	default:
		t.Fatalf("unknown daemon: %q", settings.General.UseDaemon)
	}
	if err != nil {
		t.Fatal(err)
	}

	if settings.General.UseThemer != "" {
		switch settings.General.UseThemer {
		case "matugen":
			err = theming.ApplyMatugen(wallLocation, &settings.Matugen)
		case "wallust":
			err = theming.ApplyWallust(wallLocation, &settings.Wallust)
		default:
			t.Fatalf("unknown themer: %q", settings.General.UseThemer)
		}
		if err != nil {
			t.Fatal(err)
		}
	}
}
