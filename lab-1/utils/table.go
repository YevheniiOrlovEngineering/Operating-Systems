package utils

import (
	"github.com/YevheniiOrlovEngineering/Operating-Systems/lab-1/process"
	"github.com/jedib0t/go-pretty/table"
	"os"
)

func PrintTableStdOut(pList []process.Process, order string) {
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{order, "", "", "", "", "", ""}, rowConfigAutoMerge)
	t.AppendHeader(
		table.Row{
			"#",
			"Arrival Time",
			"Burst Time",
			"Start Time",
			"Finish Time",
			"Waiting Time",
			"Turn Around Time",
		}, rowConfigAutoMerge)

	for _, p := range pList {
		t.AppendRow(table.Row{
			p.Id,
			p.At,
			p.Bt,
			p.St,
			p.Ft,
			p.Wt,
			p.Tat,
		})
		t.AppendSeparator()

	}

	t.SetStyle(table.StyleLight)
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:     order,
			WidthMin: 15,
		},
	})
	t.Render()
}
