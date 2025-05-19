package config

import (
	"github.com/spf13/viper"
)

func LoadConfig(configFilePath string) (Config, error) {

	v := viper.New()

	v.SetConfigType("toml")

	v.SetConfigFile(configFilePath)

	if err := v.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
