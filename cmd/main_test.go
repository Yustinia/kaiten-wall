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
)

func TestMain(t *testing.T) {
	projDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("project root not found: %v", err)
	}

	configPath := filepath.Join(projDir, "..", "config.example.toml")
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
	selectedWall := api.SelectRandomWall(result)

	wallLocation, err := download.DownloadWall(selectedWall, settings.General.DefaultPath)
	if err != nil {
		t.Fatal(err)
	}

	err = daemon.RunAwww(settings.General.UseDaemon, wallLocation, &settings.Awww)
	if err != nil {
		t.Fatal(err)
	}
}
