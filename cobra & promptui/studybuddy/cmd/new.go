/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"studybuddy/data"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new studybuddy note",
	Long:  `Creates a new studybuddy note`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewNote()
	},
}

type promptContent struct {
	errorMsg string
	label    string
}

func init() {
	noteCmd.AddCommand(newCmd)
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold }}",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)
	return result
}

func promptGetSelect(pc promptContent) string {
	items := []string{"animal", "food", "person", "object"}
	index := -1

	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)
	return result
}

func createNewNote() {
	wordPromptContent := promptContent{
		errorMsg: "Please provide a word",
		label:    "What word would you like to make a note of?",
	}
	word := promptGetInput(wordPromptContent)

	definitionPromptContent := promptContent{
		errorMsg: "Please provide a definition",
		label:    fmt.Sprintf("What is the definition of %s?", word),
	}
	definition := promptGetInput(definitionPromptContent)

	categoryPromptContent := promptContent{
		errorMsg: "Please provide a category",
		label:    fmt.Sprintf("What category does %s belong to?", word),
	}
	category := promptGetSelect(categoryPromptContent)

	data.InsertNote(word, definition, category)
}
