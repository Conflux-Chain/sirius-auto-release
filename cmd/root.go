/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"Conflux-Chain/sirius-auto-release/internal/runner"
	"Conflux-Chain/sirius-auto-release/internal/utils"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var configFile string

var debugMode bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sirius",
	Short: "Run the Sirius application",
	Run: func(cmd *cobra.Command, args []string) {
		logLevel := log.InfoLevel
		if debugMode {
			logLevel = log.DebugLevel
		}

		logger := log.New(os.Stdout)
		logger.SetLevel(logLevel)

		log.SetDefault(logger)

		log.Debug("Sirius application started")

		if configFile != "" {
			utils.ShowTitle("Loading config file from: " + configFile)

			cfg, err := config.LoadConfig(configFile)
			if err != nil {
				log.Error("Failed to load config file", "error", err)
				utils.ShowFailure(fmt.Sprintf("Error: Could not load config file: %v\n", err))
				return
			}
			log.Debug("Config loaded successfully", "version", cfg.Global.Version, "workdir", cfg.Global.Workdir)

			if err := runner.RunScript(&cfg); err != nil {
				log.Error("Error running script", "error", err)
				utils.ShowFailure(fmt.Sprintf("Run Error: %v\n", err))
				return
			}

		} else {
			fmt.Println("No config file provided, using default settings.")
		}
		fmt.Println("✓ All tasks completed successfully!")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "", "Path to the configuration file")

	rootCmd.Flags().BoolVarP(&debugMode, "debug", "d", false, "Enable debug mode")
}
