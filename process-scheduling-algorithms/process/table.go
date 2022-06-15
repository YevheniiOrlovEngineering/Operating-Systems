package process

import (
	"github.com/jedib0t/go-pretty/table"
	"log"
	"strconv"
)

func PrintProcessTableStdOut(pList []Process, orderMsg string, stdOut *log.Logger, avgStats ...int) {
	avgStatsMap := make(map[string]string)

	if len(avgStats) == 0 {
		avgStatsMap["Average Waiting Time"] = "-"
		avgStatsMap["Average Turn Around Time"] = "-"
	} else {
		avgStatsMap["Average Waiting Time"] = strconv.Itoa(avgStats[0])
		avgStatsMap["Average Turn Around Time"] = strconv.Itoa(avgStats[1])
	}
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	t := table.NewWriter()

	t.SetOutputMirror(stdOut.Writer())
	t.AppendHeader(table.Row{orderMsg, "", "", "", "", "", ""}, rowConfigAutoMerge)
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
	t.AppendFooter(table.Row{
		"Average Waiting Time",
		"", "", "", "", "",
		avgStatsMap["Average Waiting Time"],
	})
	t.AppendFooter(table.Row{
		"Average Turn Around Time",
		"", "", "", "", "",
		avgStatsMap["Average Turn Around Time"],
	})

	t.SetStyle(table.StyleLight)
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:     orderMsg,
			WidthMin: 15,
		},
	})
	t.Render()
}
