/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/duyanhitbe/todo-cli/internal/helpers"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use: "complete [id]",
	Long: `
Arguments:
  id  ID of todo item you want to complete.`,
	Short: "Complete a todo item",
	Run:   runComplete,
	Args:  cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(completeCmd)
}

func runComplete(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}
	index := id - 1

	data, err := helpers.ReadDataFromCSV()
	if err != nil {
		log.Fatal(err)
	}

	if id > len(data) {
		fmt.Println("Please provide valid ID")
		return
	}

	title := data[index][0]

	complete, err := strconv.ParseBool(data[index][2])
	if err != nil {
		log.Fatal(err)
	}
	if complete {
		fmt.Printf("\"%s\" has already completed\n", title)
		return
	}
	data[index][2] = fmt.Sprintf("%t", !complete)

	if err = helpers.OverrideCSV(data); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\"%s\" was completed\n", title)
}
