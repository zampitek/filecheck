package checks

import (
	"github.com/zampitek/filecheck/internal"
	"github.com/zampitek/filecheck/internal/config"
)

// CheckSize categorizes a slice of internal.FileInfo based on size.
//
// It returns three slices:
//   - lowSize: files less than 100 MB
//   - mediumSize: files between 100 MB and 1 GB
//   - highSize: files more than 1 GB
func CheckSize(files []internal.FileInfo, rules config.Rules) (lowSize, mediumSize, highSize []internal.FileInfo) {
	lowSize = make([]internal.FileInfo, 0)
	mediumSize = make([]internal.FileInfo, 0)
	highSize = make([]internal.FileInfo, 0)

	for _, file := range files {
		if file.Size < rules.Size.Low {
			lowSize = append(lowSize, file)
		} else if file.Size < rules.Size.Medium {
			mediumSize = append(mediumSize, file)
		} else {
			highSize = append(highSize, file)
		}
	}

	return lowSize, mediumSize, highSize
}

func GetEmptyFiles(files []internal.FileInfo) []internal.FileInfo {
	sortedFiles := internal.SortBySize(files)
	emptyFiles := make([]internal.FileInfo, 0)

	for i := len(sortedFiles) - 1; i > 0; i-- {
		if sortedFiles[i].Size > 0 {
			break
		}
		emptyFiles = append(emptyFiles, sortedFiles[i])
	}

	return emptyFiles
}
