package report

import (
	"fmt"
	"math"
	"strings"

	"github.com/zampitek/filecheck/internal"
	"github.com/zampitek/filecheck/internal/config"

	"github.com/fatih/color"
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

// sizeTo takes a size in bytes and converts it to
// the desired measurement unit.
//
// It returns the converted value rounded to the 2nd decimal place.
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

// makeGeneralTable creates the outline highlighting the categories of each check.
func makeGeneralTable(low, medium, high []internal.FileInfo, g string, descriptions []string) string {
	builder := strings.Builder{}
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	builder.WriteString(fmt.Sprintf("\n--- %s GROUP SUMMARY ---\n", g))

	groups := []struct {
		Label string
		Files []internal.FileInfo
	}{
		{green("LOW") + descriptions[0], low},
		{yellow("MEDIUM") + descriptions[1], medium},
		{red("HIGH") + descriptions[2], high},
	}

	for _, group := range groups {
		builder.WriteString(fmt.Sprintf("  %-55s %10d files | %5.2f GB\n", group.Label, len(group.Files), totalSize(group.Files, 3)))
	}

	builder.WriteString("--------------------------------------------------\n\n")

	return builder.String()
}

// makeTopGroupReport creates a top N ranking of the files for each category of every enabled check.
func makeTopGroupReport(files []internal.FileInfo, label string, ageTop int, colorFunc func(...any) string, description string, mU int8, sort func(slice []internal.FileInfo) []internal.FileInfo) string {
	if len(files) == 0 {
		return ""
	}

	sorted := sort(files)
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("[ %s ] - %s\n", colorFunc(label), description))
	builder.WriteString(fmt.Sprintf("  Top %d:\n", ageTop))

	for i, f := range sorted[:min(ageTop, len(sorted))] {
		builder.WriteString(fmt.Sprintf("    %d. %-105s %10d days ago | %6.2f %s\n", i+1, f.Path, f.LastAccess, sizeTo(f.Size, mU), mUToString(mU)))
	}

	builder.WriteString("\n\n")
	return builder.String()

}

// Header returns the header of the report message.
func Header() string {
	builder := strings.Builder{}
	builder.WriteString("==================================================\n")
	builder.WriteString("\t\tFILE ANALYSIS REPORT\n")
	builder.WriteString("==================================================\n")

	return builder.String()
}

// AgeReport creates and returns the report message for the age check.
func AgeReport(low, medium, high []internal.FileInfo, ageTop int, rules config.Rules) string {
	builder := strings.Builder{}

	builder.WriteString("\n###################\n")
	builder.WriteString("# BY FILE AGE     #\n")
	builder.WriteString("###################\n")

	descriptions := [3]string{
		fmt.Sprintf(" (modified within the last %d days):", rules.Age.Low),
		fmt.Sprintf(" (modified %d-%d days ago):", rules.Age.Low, rules.Age.Medium),
		fmt.Sprintf(" (modified over %d days ago):", rules.Age.Medium),
	}
	builder.WriteString(makeGeneralTable(low, medium, high, "AGE", descriptions[:]))

	if ageTop > 0 {
		builder.WriteString(makeTopGroupReport(low, "LOW", ageTop, color.New(color.FgGreen).SprintFunc(), fmt.Sprintf("Files modified in the last %d days", rules.Age.Low), 1, internal.SortByAge))
		builder.WriteString(makeTopGroupReport(medium, "MEDIUM", ageTop, color.New(color.FgYellow).SprintFunc(), fmt.Sprintf("Files modified %d-%d days ago", rules.Age.Low, rules.Age.Medium), 1, internal.SortByAge))
		builder.WriteString(makeTopGroupReport(high, "HIGH", ageTop, color.New(color.FgRed).SprintFunc(), fmt.Sprintf("Files modified over %d days ago", rules.Age.Medium), 1, internal.SortByAge))
	}

	builder.WriteString("\n\n")
	return builder.String()
}

// SizeReport creates and returns the report message for the size check.
func SizeReport(low, medium, high []internal.FileInfo, sizeTop int, rules config.Rules) string {
	builder := strings.Builder{}

	builder.WriteString("\n###################\n")
	builder.WriteString("# BY FILE SIZE    #\n")
	builder.WriteString("###################\n")

	descriptions := [3]string{
		fmt.Sprintf(" (files less than %d MB):", rules.Size.Low/1024/1024),
		fmt.Sprintf(" (files between %d MB and %d GB):", rules.Size.Low/1024/1024, rules.Size.Medium/1024/1024/1024),
		fmt.Sprintf(" (files over %d GB):", rules.Size.Medium/1024/1024/1024),
	}
	builder.WriteString(makeGeneralTable(low, medium, high, "SIZE", descriptions[:]))

	if sizeTop > 0 {
		builder.WriteString(makeTopGroupReport(low, "LOW", sizeTop, color.New(color.FgGreen).SprintFunc(), fmt.Sprintf("Files under %d MB", rules.Size.Low/1024/1024), 2, internal.SortBySize))
		builder.WriteString(makeTopGroupReport(medium, "MEDIUM", sizeTop, color.New(color.FgYellow).SprintFunc(), fmt.Sprintf("Files between %d MB and %d GB", rules.Size.Low/1024/1024, rules.Size.Medium/1024/1024/1024), 2, internal.SortBySize))
		builder.WriteString(makeTopGroupReport(high, "HIGH", sizeTop, color.New(color.FgRed).SprintFunc(), fmt.Sprintf("Files over %d GB", rules.Size.Medium/1024/1024/1024), 3, internal.SortBySize))
	}

	builder.WriteString("\n\n")

	return builder.String()
}

// EmptyFilesReport returns the number of files that have a 0-byte size
func EmptyFilesReport(emptyFiles []internal.FileInfo) string {
	red := color.New(color.FgRed).SprintFunc()

	builder := strings.Builder{}

	builder.WriteString("\n###################\n")
	builder.WriteString("# EMPTY FILES     #\n")
	builder.WriteString("###################\n")

	builder.WriteString(fmt.Sprintf("\n  %-25s %10d files\n", red(" 0-SIZED FILES "), len(emptyFiles)))
	builder.WriteString("--------------------------------------------------\n\n")

	builder.WriteString("\n\n")

	return builder.String()
}
