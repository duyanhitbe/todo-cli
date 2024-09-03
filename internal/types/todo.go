package types

import (
	"time"
)

type Todo struct {
	Title    string
	Date     time.Time
	Complete bool
}

func NewTodo(title string) *Todo {
	return &Todo{
		Title:    title,
		Date:     time.Now(),
		Complete: false,
	}
}
