package report

import (
	"fmt"
	"math"
	"strings"

	"filecheck/internal"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

func totalSize(files []internal.FileInfo, mU int8) float32 {
	var totalBytes int64
	for _, file := range files {
		totalBytes += file.Size
	}

	bytes := float32(totalBytes)
	var result float32

	switch mU {
	case 0:
		result = bytes
	case 1:
		result = bytes / 1024
	case 2:
		result = bytes / (1024 * 1024)
	default:
		result = bytes / (1024 * 1024 * 1024)
	}

	return float32(math.Round(float64(result)*100) / 100)
}

func AgeReport(lowAge, mediumAge, highAge []internal.FileInfo) string {
	builder := strings.Builder{}
	builder.WriteString("\nAGE GROUP SUMMARY\n")
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

	t := table.NewWriter()
	t.SetOutputMirror(&builder)
	t.AppendHeader(table.Row{"GROUP", "FILE COUNT", "TOTAL SIZE"})
	t.AppendRow(table.Row{green("LOW"), len(lowAge), blue(fmt.Sprintf("%.2f GB", totalSize(lowAge, 3)))})
	t.AppendRow(table.Row{yellow("MEDIUM"), len(mediumAge), blue(fmt.Sprintf("%.2f GB", totalSize(mediumAge, 3)))})
	t.AppendRow(table.Row{red("HIGH"), len(highAge), blue(fmt.Sprintf("%.2f GB", totalSize(highAge, 3)))})

	t.Render()
	builder.WriteString("\n\n")

	return builder.String()
}
