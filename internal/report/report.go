package report

import (
	"fmt"
	"math"
	"strings"

	"github.com/zampitek/filecheck/internal"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

func sizeTo(size int64, mU int8) float32 {
	div := float32(1)
	switch mU {
	case 1:
		div = 1024
	case 2:
		div = 1024 * 1024
	case 3:
		div = 1024 * 1024 * 1024
	}
	return float32(math.Round(float64(float32(size)/div)*100) / 100)
}

func totalSize(files []internal.FileInfo, mU int8) float32 {
	var total int64
	for _, f := range files {
		total += f.Size
	}
	return sizeTo(total, mU)
}

func makeGeneralAgeTable(low, medium, high []internal.FileInfo) string {
	builder := strings.Builder{}
	blue := color.New(color.FgBlue).SprintFunc()

	groups := []struct {
		Label string
		Files []internal.FileInfo
		Color func(a ...any) string
	}{
		{"LOW", low, color.New(color.FgGreen).SprintFunc()},
		{"MEDIUM", medium, color.New(color.FgYellow).SprintFunc()},
		{"HIGH", high, color.New(color.FgRed).SprintFunc()},
	}

	t := table.NewWriter()
	t.SetOutputMirror(&builder)
	t.AppendHeader(table.Row{"GROUP", "FILE COUNT", "TOTAL SIZE"})

	for _, g := range groups {
		t.AppendRow(table.Row{
			g.Color(g.Label),
			len(g.Files),
			blue(fmt.Sprintf("%.2f GB", totalSize(g.Files, 3))),
		})
	}

	t.Render()
	builder.WriteString("\n\n")
	return builder.String()
}

func makeTopGroupReport(files []internal.FileInfo, label string, ageTop int, colorFunc func(...any) string, description string) string {
	if len(files) == 0 {
		return ""
	}

	sorted := internal.SortByAge(files)
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("AGE GROUP: %s (%s)\n", colorFunc(label), description))

	t := table.NewWriter()
	t.SetOutputMirror(&builder)
	t.AppendHeader(table.Row{"FILE PATH", "MODIFIED", "SIZE (KB)"})

	for _, f := range sorted[:min(ageTop, len(sorted))] {
		t.AppendRow(table.Row{f.Path, f.LastAccess, sizeTo(f.Size, 1)})
	}

	t.Render()
	builder.WriteString("\n\n")
	return builder.String()
}

func AgeReport(low, medium, high []internal.FileInfo, ageTop int) string {
	builder := strings.Builder{}
	builder.WriteString("\n===================================================AGE GROUP SUMMARY===================================================\n\n")
	builder.WriteString(makeGeneralAgeTable(low, medium, high))

	if ageTop > 0 {
		builder.WriteString(makeTopGroupReport(low, "LOW", ageTop, color.New(color.FgGreen).SprintFunc(), "files modified in the last 90 days"))
		builder.WriteString(makeTopGroupReport(medium, "MEDIUM", ageTop, color.New(color.FgYellow).SprintFunc(), "files modified 90-180 days ago"))
		builder.WriteString(makeTopGroupReport(high, "HIGH", ageTop, color.New(color.FgRed).SprintFunc(), "files modified over 180 days ago"))
	}

	builder.WriteString("\n\n")
	return builder.String()
}
