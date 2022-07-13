package render

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type TableRenderer struct {
	Headers        []string
	Rows           [][]string
	AutoWrapText   bool
	RowLine        bool
	AutoMergeCells bool
}

func (t TableRenderer) RenderOutput() {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(t.Headers)
	table.SetAutoWrapText(t.AutoWrapText)
	table.SetRowLine(t.RowLine)
	table.SetAutoMergeCells(t.AutoMergeCells)

	table.AppendBulk(t.Rows)
	table.Render()
}
