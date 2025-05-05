package report

import (
	"filecheck/internal/scanner"
	"fmt"
	"sort"
	"strings"
)

var ageLow []scanner.FileInfo
var ageMedium []scanner.FileInfo
var ageHigh []scanner.FileInfo
var Top5Oldest []scanner.FileInfo

var sizeLow []scanner.FileInfo
var sizeMedium []scanner.FileInfo
var sizeHigh []scanner.FileInfo
var Top5Heaviest []scanner.FileInfo

type Report struct {
	Modified90   []scanner.FileInfo
	Modified180  []scanner.FileInfo
	ModifiedHigh []scanner.FileInfo

	Size10   []scanner.FileInfo
	Size100  []scanner.FileInfo
	sizeHigh []scanner.FileInfo
}

type ExtendedReport struct {
	Modified90   []scanner.FileInfo
	Modified180  []scanner.FileInfo
	ModifiedHigh []scanner.FileInfo
	Top5Oldest   []scanner.FileInfo

	Size10       []scanner.FileInfo
	Size100      []scanner.FileInfo
	sizeHigh     []scanner.FileInfo
	Top5Heaviest []scanner.FileInfo
}

func cathegorizeFilesByAge(file scanner.FileInfo) {
	if file.LastAccess < 30 {
		return
	} else if file.LastAccess < 90 {
		ageLow = append(ageLow, file)
	} else if file.LastAccess < 180 {
		ageMedium = append(ageMedium, file)
	} else {
		ageHigh = append(ageHigh, file)
	}
}

func cathegorizeFilesBySize(file scanner.FileInfo) {
	if file.IsDir {
		return
	}

	if file.Size < 50*1024*1024 {
		sizeLow = append(sizeLow, file)
	} else if file.Size < 500*1024*1024 {
		sizeMedium = append(sizeMedium, file)
	} else {
		sizeHigh = append(sizeHigh, file)
	}
}

func CreateReport(report scanner.ScanReport) Report {
	for _, f := range report.Files {
		cathegorizeFilesByAge(f)
		cathegorizeFilesBySize(f)
	}

	return Report{ageLow, ageMedium, ageHigh,
		sizeLow, sizeMedium, sizeHigh}
}

func CreateExtendedReport(report scanner.ScanReport) ExtendedReport {
	for _, f := range report.Files {
		cathegorizeFilesByAge(f)
		cathegorizeFilesBySize(f)
	}

	sort.Slice(ageHigh, func(i, j int) bool {
		return ageHigh[i].LastAccess > ageHigh[j].LastAccess
	})

	numFilesAge := len(ageHigh)
	countAge := min(numFilesAge, 5)
	Top5Oldest = ageHigh[:countAge]

	sort.Slice(sizeHigh, func(i, j int) bool {
		return sizeHigh[i].Size > sizeHigh[j].Size
	})

	numFilesSize := len(sizeHigh)
	countSize := min(numFilesSize, 5)
	Top5Heaviest = sizeHigh[:countSize]

	return ExtendedReport{ageLow, ageMedium, ageHigh, Top5Oldest,
		sizeLow, sizeMedium, sizeHigh, Top5Heaviest}
}

func PrintReport(report Report) string {
	var builder strings.Builder
	var problems = false

	if len(report.Modified90) != 0 || len(report.Modified180) != 0 || len(report.ModifiedHigh) != 0 {
		builder.WriteString("Found several file modified over 30 days ago:\n\n")
		builder.WriteString(fmt.Sprintf("\tLOW SEVERITY (modified within the last 90 days): %d files\n", len(report.Modified90)))
		builder.WriteString(fmt.Sprintf("\tMEDIUM SEVERITY (modified within the last 180 days): %d files\n", len(report.Modified180)))
		builder.WriteString(fmt.Sprintf("\tHIGH SEVERITY (modified over 180 days ago): %d files\n", len(report.ModifiedHigh)))

		problems = true
	}
	if len(report.Size10) != 0 || len(report.Size100) != 0 || len(report.sizeHigh) != 0 {
		builder.WriteString("\nFound several file with a considerable size:\n\n")
		builder.WriteString(fmt.Sprintf("\tLOW SEVERITY (size less than 50 MB): %d files\n", len(report.Size10)))
		builder.WriteString(fmt.Sprintf("\tMEDIUM SEVERITY (size less than 500 MB): %d files\n", len(report.Size100)))
		builder.WriteString(fmt.Sprintf("\tHIGH SEVERITY (size over 500 MB): %d files\n", len(report.sizeHigh)))

		problems = true
	}

	if !problems {
		builder.WriteString("No problems found\n")
	}

	return builder.String()
}

func PrintExtendedReport(report ExtendedReport) string {
	var builder strings.Builder
	var problems = false

	if len(report.Modified90) != 0 || len(report.Modified180) != 0 || len(report.ModifiedHigh) != 0 {
		builder.WriteString("Found several file modified over 30 days ago:\n\n")
		builder.WriteString(fmt.Sprintf("\tLOW SEVERITY (modified within the last 90 days): %d files\n", len(report.Modified90)))
		builder.WriteString(fmt.Sprintf("\tMEDIUM SEVERITY (modified within the last 180 days): %d files\n", len(report.Modified180)))
		builder.WriteString(fmt.Sprintf("\tHIGH SEVERITY (modified over 180 days ago): %d files\n\n", len(report.ModifiedHigh)))

		for i, file := range report.Top5Oldest {
			builder.WriteString(fmt.Sprintf("\t%d. %s\t(%d days)\n", i, file.Path, file.LastAccess))
		}

		problems = true
	}
	if len(report.Size10) != 0 || len(report.Size100) != 0 || len(report.sizeHigh) != 0 {
		builder.WriteString("\nFound several file with a considerable size:\n\n")
		builder.WriteString(fmt.Sprintf("\tLOW SEVERITY (size less than 50 MB): %d files\n", len(report.Size10)))
		builder.WriteString(fmt.Sprintf("\tMEDIUM SEVERITY (size less than 500 MB): %d files\n", len(report.Size100)))
		builder.WriteString(fmt.Sprintf("\tHIGH SEVERITY (size over 500 MB): %d files\n\n", len(report.sizeHigh)))

		for i, file := range report.Top5Heaviest {
			builder.WriteString(fmt.Sprintf("\t%d. %s\t(%d MB)\n", i, file.Path, file.Size/1024/1024))
		}

		problems = true
	}

	if !problems {
		builder.WriteString("No problems found\n")
	}

	return builder.String()
}
