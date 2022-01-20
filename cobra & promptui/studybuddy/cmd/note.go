/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "A note can be anything you'd like to study and review",
	Long:  `A note can be anything you'd like to study and review`,
}

func init() {
	rootCmd.AddCommand(noteCmd)
}
