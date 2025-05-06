package report

import (
	"sort"

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
