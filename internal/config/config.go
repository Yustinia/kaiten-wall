package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/Yustinia/kaiten-wall/internal/fs"
)

func ParseSettings(configPath string) (ConfigModel, error) {
	err := fs.IsFileExist(configPath)
	if err != nil {
		return ConfigModel{}, err
	}

	var configInstance ConfigModel
	_, err = toml.DecodeFile(configPath, &configInstance)
	if err != nil {
		return ConfigModel{}, fmt.Errorf("failed to decode: %w", err)
	}

	return configInstance, nil
}
