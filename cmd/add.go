/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/duyanhitbe/todo-cli/internal/helpers"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use: "add [title]",
	Long: `
Arguments:
  title  Title of todo item.`,
	Short: "Add a new todo item",
	Run:   runAdd,
	Args:  cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func runAdd(cmd *cobra.Command, args []string) {
	title := args[0]

	file, err := helpers.OpenFileCSV()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	helpers.WriteTodo(title, file)
	fmt.Println(title + " was added")
}
