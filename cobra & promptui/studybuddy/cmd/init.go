/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"studybuddy/data"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a new studybuddy database and table",
	Long:  `Initialise a new studybuddy database and table`,
	Run: func(cmd *cobra.Command, args []string) {
		data.CreateTable()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
