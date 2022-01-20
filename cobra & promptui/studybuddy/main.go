/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"log"
	"studybuddy/cmd"
	"studybuddy/data"
)

func main() {
	if err := data.OpenDatabase(); err != nil {
		log.Fatal(err.Error())
	}
	cmd.Execute()
}
