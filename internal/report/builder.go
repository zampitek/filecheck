package report

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

func PrintReport(report Report) {
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

	fmt.Printf("\n\t[%s] \t%s\n", green("LOW AGE"), blue(fmt.Sprint(len(report.Modified90))))
	fmt.Printf("\t[%s] \t%s\n", yellow("MEDIUM AGE"), blue(fmt.Sprint(len(report.Modified180))))
	fmt.Printf("\t[%s] \t%s\n", red("HIGH AGE"), blue(fmt.Sprint(len(report.ModifiedHigh))))

	fmt.Printf("\n\t[%s] \t%s\n", green("LOW SIZE"), blue(fmt.Sprint(len(report.Size10))))
	fmt.Printf("\t[%s] \t%s\n", yellow("MEDIUM SIZE"), blue(fmt.Sprint(len(report.Size100))))
	fmt.Printf("\t[%s] \t%s\n", red("HIGH SIZE"), blue(fmt.Sprint(len(report.sizeHigh))))
}

func PrintExtendedReport(report ExtendedReport) {
	PrintReport(report.Report)
	fmt.Print("\n\n")

	ageTable := table.NewWriter()
	ageTable.SetOutputMirror(os.Stdout)
	ageTable.AppendHeader(table.Row{"#", "Path", "Age"})
	blue := color.New(color.FgBlue).SprintFunc()

	if len(report.Top5Oldest) > 0 {
		for i, file := range report.Top5Oldest {
			ageTable.AppendRow(table.Row{i + 1, file.Path, blue(file.LastAccess)})
		}
	}

	ageTable.Render()

	sizeTable := table.NewWriter()
	sizeTable.SetOutputMirror(os.Stdout)
	sizeTable.AppendHeader(table.Row{"#", "Path", "Size (MB)"})

	if len(report.Top5Heaviest) > 0 {
		for i, file := range report.Top5Heaviest {
			sizeTable.AppendRow(table.Row{i + 1, file.Path, blue(file.Size / 1024 / 1024)})
		}
	}

	sizeTable.Render()
}
