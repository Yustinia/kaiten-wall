package defaults

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
)

//go:embed config.example.toml
var defaultConfig []byte

var ErrFirstLaunch = errors.New("first launch detected, please configure before running again")

func FirstLaunch(configDir string, configPath string) error {
	if _, err := os.Stat(configPath); err == nil {
		return nil
	}

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	if err := os.WriteFile(configPath, defaultConfig, 0644); err != nil {
		return fmt.Errorf("failed to write default config: %w", err)
	}

	return ErrFirstLaunch
}
