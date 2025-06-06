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

		path := "./config.toml"
		if outputConfigFile != "" {
			path = outputConfigFile
		}

		tomlBytes, err := toml.Marshal(useConfig)

		if err := utils.WriteToFile(tomlBytes, path); err != nil {
			utils.ShowFailure(fmt.Sprintf("Failed to write config to file: %v", err))
			return
		}

		utils.ShowSuccess(fmt.Sprintf("Configuration file generated successfully: %s", path))
		utils.ShowCommandSuggestion("To use this configuration, run:")
		utils.ShowCommand(fmt.Sprintf("sirius --config %s", path))

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&outputConfigFile, "output", "o", "config.toml", "Output configuration file name")
}
