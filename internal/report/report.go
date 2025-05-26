package report

import (
	"fmt"
	"math"
	"strings"

	"github.com/fatih/color"
	"github.com/zampitek/filecheck/internal"
	"github.com/zampitek/filecheck/internal/config"
)

// mUToString converts a mU number from int8 to string,
// for the correspondent measurement unit.
func mUToString(mU int8) string {
	switch mU {
	case 1:
		return "KB"
	case 2:
		return "MB"
	case 3:
		return "GB"
	default:
		return "B"
	}
}

// sizeTo takes a size in bytes and converts it to the desired
// measurement unit, rounded at the 2nd decimal place.
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

// totalSize returns the sum of the sizes of each file in a given slice.
func totalSize(files []internal.FileInfo, mU int8) float32 {
	var total int64
	for _, f := range files {
		total += f.Size
	}
	return sizeTo(total, mU)
}

var (
	green  = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
)

// Header returns the header of the report message.
func Header() string {
	return `==================================================
		FILE ANALYSIS REPORT
==================================================
`
}

// buildGroupSummary creates the outline highlighting the categories of each check.
func buildGroupSummary(title string, low, medium, high []internal.FileInfo, descriptions []string) string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("\n--- %s GROUP SUMMARY ---\n", title))

	groups := []struct {
		Label string
		Files []internal.FileInfo
	}{
		{green("LOW") + descriptions[0], low},
		{yellow("MEDIUM") + descriptions[1], medium},
		{red("HIGH") + descriptions[2], high},
	}

	for _, g := range groups {
		builder.WriteString(fmt.Sprintf("  %-55s %10d files | %5.2f GB\n", g.Label, len(g.Files), totalSize(g.Files, 3)))
	}

	builder.WriteString("--------------------------------------------------\n\n")
	return builder.String()
}

// buildTopNReport creates a top N ranking of the files for each category of every enabled check.
func buildTopNReport(label string, description string, files []internal.FileInfo, top int, colorFn func(...any) string, mU int8, sortFn func([]internal.FileInfo) []internal.FileInfo) string {
	if len(files) == 0 {
		return ""
	}

	sorted := sortFn(files)
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("[ %s ] - %s\n", colorFn(label), description))
	builder.WriteString(fmt.Sprintf("  Top %d:\n", top))

	for i, f := range sorted[:internal.Min(top, len(sorted))] {
		builder.WriteString(fmt.Sprintf("    %d. %-105s %10d days ago | %6.2f %s\n",
			i+1, f.Path, f.LastAccess, sizeTo(f.Size, mU), mUToString(mU)))
	}

	builder.WriteString("\n\n")
	return builder.String()
}

// AgeReport creates and returns the report message for the age check.
func AgeReport(low, medium, high []internal.FileInfo, top int, rules config.Rules) string {
	builder := strings.Builder{}

	builder.WriteString("\n###################\n# BY FILE AGE     #\n###################\n")

	descriptions := []string{
		fmt.Sprintf(" (modified within the last %d days):", rules.Age.Low),
		fmt.Sprintf(" (modified %d-%d days ago):", rules.Age.Low, rules.Age.Medium),
		fmt.Sprintf(" (modified over %d days ago):", rules.Age.Medium),
	}
	builder.WriteString(buildGroupSummary("AGE", low, medium, high, descriptions))

	if top > 0 {
		builder.WriteString(buildTopNReport("LOW", fmt.Sprintf("Files modified in the last %d days", rules.Age.Low), low, top, green, 1, internal.SortByAge))
		builder.WriteString(buildTopNReport("MEDIUM", fmt.Sprintf("Files modified %d-%d days ago", rules.Age.Low, rules.Age.Medium), medium, top, yellow, 1, internal.SortByAge))
		builder.WriteString(buildTopNReport("HIGH", fmt.Sprintf("Files modified over %d days ago", rules.Age.Medium), high, top, red, 1, internal.SortByAge))
	}

	return builder.String()
}

// SizeReport creates and returns the report message for the size check.
func SizeReport(low, medium, high []internal.FileInfo, top int, rules config.Rules) string {
	builder := strings.Builder{}

	builder.WriteString("\n###################\n# BY FILE SIZE    #\n###################\n")

	descriptions := []string{
		fmt.Sprintf(" (files less than %d MB):", rules.Size.Low/1024/1024),
		fmt.Sprintf(" (files between %d MB and %d GB):", rules.Size.Low/1024/1024, rules.Size.Medium/1024/1024/1024),
		fmt.Sprintf(" (files over %d GB):", rules.Size.Medium/1024/1024/1024),
	}
	builder.WriteString(buildGroupSummary("SIZE", low, medium, high, descriptions))

	if top > 0 {
		builder.WriteString(buildTopNReport("LOW", fmt.Sprintf("Files under %d MB", rules.Size.Low/1024/1024), low, top, green, 2, internal.SortBySize))
		builder.WriteString(buildTopNReport("MEDIUM", fmt.Sprintf("Files between %d MB and %d GB", rules.Size.Low/1024/1024, rules.Size.Medium/1024/1024/1024), medium, top, yellow, 2, internal.SortBySize))
		builder.WriteString(buildTopNReport("HIGH", fmt.Sprintf("Files over %d GB", rules.Size.Medium/1024/1024/1024), high, top, red, 3, internal.SortBySize))
	}

	return builder.String()
}

// EmptyFilesReport returns the number of files that have a 0-byte size
func EmptyFilesReport(emptyFiles []internal.FileInfo) string {
	builder := strings.Builder{}

	builder.WriteString("\n###################\n# EMPTY FILES     #\n###################\n")
	builder.WriteString(fmt.Sprintf("\n  %-25s %10d files\n", red(" 0-SIZED FILES "), len(emptyFiles)))
	builder.WriteString("--------------------------------------------------\n\n\n\n")

	return builder.String()
}
