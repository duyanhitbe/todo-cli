/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strconv"

	"github.com/duyanhitbe/todo-cli/internal/helpers"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [index]",
	Short: "Remove a todo item",
	Run:   runRemove,
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

func runRemove(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}
	index := id - 1

	data, err := helpers.ReadDataFromCSV()
	if err != nil {
		log.Fatal(err)
	}

	data = append(data[:index], data[index+1:]...)
	if err = helpers.OverrideCSV(data); err != nil {
		log.Fatal(err)
	}
}
