/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"Conflux-Chain/sirius-auto-release/internal/interactive"
	"Conflux-Chain/sirius-auto-release/internal/utils"
	"fmt"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

var outputConfigFile string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize and generate a configuration file interactively",
	Run: func(cmd *cobra.Command, args []string) {

		cfg := config.Config{}

		cfg.Global.Version = 1

		useConfig, err := interactive.CollectConfig(&cfg)

		if err != nil {
			utils.ShowFailure(fmt.Sprintf("Failed to collect configuration: %v", err))
			return
		}

		var themeConfig config.ThemeConfig
		if err := toml.Unmarshal([]byte(config.DefaultESpaceTheme), &themeConfig); err != nil {
			utils.ShowFailure(fmt.Sprintf("Failed to unmarshal theme config: %v", err))
			return
		}

		if (cfg.Global.Space == config.ALL_SPACE || cfg.Global.Space == config.E_SPACE) && cfg.Frontend.ESpaceSettings.Enabled {
			cfg.Frontend.ESpaceSettings.EnvTheme = themeConfig
		}

		path := "./config.toml"
		if outputConfigFile != "" {
			path = outputConfigFile
		}

		tomlBytes, err := toml.Marshal(useConfig)

		if err := utils.WriteToFile(tomlBytes, path); err != nil {
			utils.ShowFailure(fmt.Sprintf("Failed to write config to file: %v", err))
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&outputConfigFile, "output", "o", "config.toml", "Output configuration file name")
}
