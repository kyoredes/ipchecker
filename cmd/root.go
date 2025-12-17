/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ipchecker/checker"
	"os"

	"github.com/spf13/cobra"
)

var logFormat string
var validate bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ipchecker [list.txt]",
	Short: "Checks IPs in a list",
	Long:  "Checks IPs in a list",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		file := args[0]
		return checker.Run(file)
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		error := checker.InitLogger(logFormat)
		if error != nil {
			panic(error)
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ipchecker.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVar(
		&logFormat,
		"log-format",
		"json",
		"Log format. Supported formats: json, text",
	)
}
