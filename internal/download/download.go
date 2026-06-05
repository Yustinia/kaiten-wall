package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Yustinia/kaiten-wall/internal/config"
)

func DownloadWall(wallpaper string, settings config.ConfigModel) (string, error) {
	resp, err := http.Get(wallpaper)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status: %w", err)
	}

	ext := filepath.Ext(wallpaper)
	outPath := filepath.Join(settings.General.DefaultPath, "wallhaven", ext)

	file, err := os.Create(outPath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	return outPath, nil
}
