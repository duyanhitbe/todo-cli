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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of todo item",
	Run:   runList,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func runList(cmd *cobra.Command, args []string) {
	//Read data from file
	data, err := helpers.ReadDataFromCSV()
	if err != nil {
		log.Fatal(err)
	}

	if len(data) == 0 {
		fmt.Println("List is empty now. Please add more todo item!")
		return
	}

	table := helpers.NewTable()
	table.Write(data)
	table.Print()
}
