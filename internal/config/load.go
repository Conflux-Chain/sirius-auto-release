package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

func LoadConfig(configFilePath string) (Config, error) {
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := toml.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
