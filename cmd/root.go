/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sirius",
	Short: "Run the Sirius application",
	Run: func(cmd *cobra.Command, args []string) {

		if configFile != "" {
			fmt.Printf("Using config file: %s\n", configFile)
		} else {
			fmt.Println("No config file provided, using default settings.")
		}

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
}
