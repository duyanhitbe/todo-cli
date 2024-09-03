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
	Use:   "complete [id]",
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

	complete, err := strconv.ParseBool(data[index][2])
	if err != nil {
		log.Fatal(err)
	}
	data[index][2] = fmt.Sprintf("%t", !complete)

	if err = helpers.OverrideCSV(data); err != nil {
		log.Fatal(err)
	}
}
