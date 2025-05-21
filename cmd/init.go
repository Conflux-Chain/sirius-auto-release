/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"Conflux-Chain/sirius-auto-release/internal/interactive"
	"Conflux-Chain/sirius-auto-release/internal/utils"
	"fmt"

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

		fmt.Printf("todo %v", useConfig)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&outputConfigFile, "output", "o", "config.toml", "Output configuration file name")
}
