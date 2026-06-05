package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadWall(wallpaper string, wallOutPath string) (string, error) {
	resp, err := http.Get(wallpaper)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	ext := filepath.Ext(wallpaper)
	outPath := filepath.Join(wallOutPath, fmt.Sprintf("wallhaven%s", ext))

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
