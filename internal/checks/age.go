package checks

import (
	"github.com/zampitek/filecheck/internal"
)

// CheckAge categorizes a slice of internal.FileInfo based on last access time.
//
// It returns three slices:
//   - lowAge: files accessed within the last 90 days
//   - mediumAge: files accessed between 90 and 180 days ago
//   - highAge: files accessed more than 180 days ago
func CheckAge(files []internal.FileInfo) (lowAge, mediumAge, highAge []internal.FileInfo) {
	lowAge = make([]internal.FileInfo, 0)
	mediumAge = make([]internal.FileInfo, 0)
	highAge = make([]internal.FileInfo, 0)

	for _, file := range files {
		if file.LastAccess < 90 {
			lowAge = append(lowAge, file)
		} else if file.LastAccess < 180 {
			mediumAge = append(mediumAge, file)
		} else {
			highAge = append(highAge, file)
		}
	}

	return lowAge, mediumAge, highAge
}
