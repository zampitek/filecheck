package report

import (
	"fmt"
	"sort"
	"strings"

	"github.com/zampitek/filecheck/internal/scanner"
)

type Report struct {
	Modified90   []scanner.FileInfo
	Modified180  []scanner.FileInfo
	ModifiedHigh []scanner.FileInfo
	Size10       []scanner.FileInfo
	Size100      []scanner.FileInfo
	sizeHigh     []scanner.FileInfo
}

type ExtendedReport struct {
	Report
	Top5Oldest   []scanner.FileInfo
	Top5Heaviest []scanner.FileInfo
}

func cathegorizeFiles(files []scanner.FileInfo) (age90, age180, ageHigh, size50, size500, sizeHigh []scanner.FileInfo) {
	for _, file := range files {
		if !file.IsDir {
			switch {
			case file.Size < 50*1024*1024:
				size50 = append(size50, file)
			case file.Size < 500*1024*1024:
				size500 = append(size500, file)
			default:
				sizeHigh = append(sizeHigh, file)
			}
		}

		switch {
		case file.LastAccess < 90:
			age90 = append(age90, file)
		case file.LastAccess < 180:
			age180 = append(age180, file)
		default:
			ageHigh = append(ageHigh, file)
		}
	}

	return
}

func CreateReport(report scanner.ScanReport) Report {
	age90, age180, ageHigh, size50, size500, sizeHigh := cathegorizeFiles(report.Files)
	return Report{age90, age180, ageHigh, size50, size500, sizeHigh}
}

func CreateExtendedReport(report scanner.ScanReport) ExtendedReport {
	r := CreateReport(report)

	sort.Slice(r.ModifiedHigh, func(i, j int) bool {
		return r.ModifiedHigh[i].LastAccess > r.ModifiedHigh[j].LastAccess
	})
	topAge := min(len(r.ModifiedHigh), 5)
	topOldest := r.ModifiedHigh[:topAge]

	sort.Slice(r.sizeHigh, func(i, j int) bool {
		return r.sizeHigh[i].Size > r.sizeHigh[j].Size
	})
	topSize := min(len(r.sizeHigh), 5)
	topHeaviest := r.sizeHigh[:topSize]

	return ExtendedReport{Report: r, Top5Oldest: topOldest, Top5Heaviest: topHeaviest}
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
	output := PrintReport(report.Report)
	var builder strings.Builder
	builder.WriteString(output)

	if len(report.Top5Oldest) > 0 {
		builder.WriteString("\nTop 5 oldest files:\n")
		for i, file := range report.Top5Oldest {
			builder.WriteString(fmt.Sprintf("\t%d. %s\t (%d days)\n", i+1, file.Path, file.LastAccess))
		}
	}

	if len(report.Top5Heaviest) > 0 {
		builder.WriteString("\nTop 5 heaviest files:\n")
		for i, file := range report.Top5Heaviest {
			builder.WriteString(fmt.Sprintf("\t%d. %s\t (%d MB)\n", i+1, file.Path, file.Size/1024/1024))
		}
	}

	return builder.String()
}
