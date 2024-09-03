package helpers

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
)

type TabWriter struct {
	Writer *tabwriter.Writer
}

func NewTabWriter() *TabWriter {
	w := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0)
	w.Write([]byte("ID\tTitle\tDate\tComplete\n"))

	return &TabWriter{
		Writer: w,
	}
}

func (tw *TabWriter) Write(i int, row []string) {
	id := i + 1
	title := row[0]
	date, _ := time.Parse(time.RFC3339, row[1])
	complete := row[2]
	s := fmt.Sprintf("%d\t%s\t%s\t%s\n", id, title, timediff.TimeDiff(date), complete)

	tw.Writer.Write([]byte(s))
}
