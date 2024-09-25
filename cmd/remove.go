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

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use: "remove [id]",
	Long: `
Arguments:
  id  ID of todo item you want to remove.`,
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

	if id > len(data) {
		fmt.Println("Please provide valid ID")
		return
	}

	title := data[index][0]

	data = append(data[:index], data[index+1:]...)
	if err = helpers.OverrideCSV(data); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\"%s\" was deleted\n", title)
}
