/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Conflux-Chain/sirius-auto-release/internal/config"
	"Conflux-Chain/sirius-auto-release/internal/runner"
	"Conflux-Chain/sirius-auto-release/internal/utils"
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var configFile string

var debugMode bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sirius",
	Short: "Run the Sirius application",
	Run: func(cmd *cobra.Command, args []string) {
		logLevel := slog.LevelInfo
		if debugMode {
			logLevel = slog.LevelDebug
		}

		opts := &slog.HandlerOptions{
			Level: logLevel,
		}

		handler := slog.NewTextHandler(os.Stderr, opts)

		logger := slog.New(handler)

		slog.SetDefault(logger)
		slog.Debug("Sirius application started")

		if configFile != "" {
			utils.ShowTitle("Loading config file from: " + configFile)

			cfg, err := config.LoadConfig(configFile)
			if err != nil {
				slog.Error("Failed to load config file", "error", err)
				utils.ShowFailure(fmt.Sprintf("Error: Could not load config file: %v\n", err))
				return
			}
			slog.Debug("Config loaded successfully", "version", cfg.Global.Version, "workdir", cfg.Global.Workdir)

			if err := runner.RunScript(&cfg); err != nil {
				slog.Error("Error running script", "error", err)
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
