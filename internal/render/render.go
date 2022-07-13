package render

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func RenderOutput(headers []string, rows [][]string) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(headers)
	table.SetAutoWrapText(false)
	table.SetRowLine(true)

	table.AppendBulk(rows)
	table.Render()
}
