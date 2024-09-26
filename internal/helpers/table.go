package helpers

import (
	"fmt"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/mergestat/timediff"
)

type Table struct {
	table simpletable.Table
}

func NewTable() *Table {
	return &Table{
		table: *simpletable.New(),
	}
}

func (t *Table) Write(rows [][]string) {
	t.table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "Title"},
			{Align: simpletable.AlignCenter, Text: "Date"},
			{Align: simpletable.AlignCenter, Text: "Complete"},
		},
	}

	todoCount := 0
	for i, row := range rows {
		id := i + 1
		title := row[0]
		date, _ := time.Parse(time.RFC3339, row[1])
		complete := row[2]

		if complete == "false" {
			todoCount++
		}

		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", id)},
			{Align: simpletable.AlignCenter, Text: title},
			{Align: simpletable.AlignCenter, Text: timediff.TimeDiff(date)},
			{Align: simpletable.AlignCenter, Text: complete},
		}

		t.table.Body.Cells = append(t.table.Body.Cells, r)
	}

	t.table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 4, Text: fmt.Sprintf("You have %d todo item", todoCount)},
		},
	}

	t.table.SetStyle(simpletable.StyleUnicode)
}

func (t *Table) Print() {
	t.table.Println()
}
