package helpers

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
)

const tab = "\t|\t"

type TabWriter struct {
	Writer *tabwriter.Writer
}

func NewTabWriter() *TabWriter {
	w := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0)
	w.Write([]byte(fmt.Sprintf("ID%sTitle%sDate%sComplete%s\n", tab, tab, tab, tab)))

	return &TabWriter{
		Writer: w,
	}
}

func (tw *TabWriter) Write(i int, row []string) {
	id := i + 1
	title := row[0]
	date, _ := time.Parse(time.RFC3339, row[1])
	complete := row[2]
	s := fmt.Sprintf("%d%s%s%s%s%s%s%s\n", id, tab, title, tab, timediff.TimeDiff(date), tab, complete, tab)

	tw.Writer.Write([]byte(s))
}
