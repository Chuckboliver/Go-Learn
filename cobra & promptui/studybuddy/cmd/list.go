/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"studybuddy/data"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "See list of all notes you've added",
	Long:  `See list of all notes you've added`,
	Run: func(cmd *cobra.Command, args []string) {
		data.DisplayAllNotes()
	},
}

func init() {
	noteCmd.AddCommand(listCmd)
}
