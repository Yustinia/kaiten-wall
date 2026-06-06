package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

func ParseSettings(configPath string) (ConfigModel, error) {
	var configInstance ConfigModel
	_, err := toml.DecodeFile(configPath, &configInstance)
	if err != nil {
		return ConfigModel{}, fmt.Errorf("failed to decode: %w", err)
	}

	return configInstance, nil
}
