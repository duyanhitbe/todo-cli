package helpers

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/duyanhitbe/todo-cli/internal/types"
)

const (
	fileName = "data.csv"
)

func OpenFileCSV() (file *os.File, err error) {
	file, err = os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	return
}

func ReadDataFromCSV() (data [][]string, err error) {
	file, err := OpenFileCSV()
	if err != nil {
		return
	}
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err = reader.ReadAll()
	if err != nil {
		return
	}
	return
}

func WriteTodo(title string, file *os.File) {
	writer := csv.NewWriter(file)
	defer writer.Flush()

	todo := types.NewTodo(title)
	writer.Write([]string{
		todo.Title,
		todo.Date.Format(time.RFC3339),
		fmt.Sprintf("%t", todo.Complete),
	})
}

func OverrideCSV(data [][]string) (err error) {
	file, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range data {
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return
}
