package checks

import (
	"github.com/zampitek/filecheck/internal"
	"github.com/zampitek/filecheck/internal/config"
)

// CheckAge categorizes a slice of internal.FileInfo based on last access time.
//
// It returns three slices:
//   - lowAge: files accessed within the last 90 days
//   - mediumAge: files accessed between 90 and 180 days ago
//   - highAge: files accessed more than 180 days ago
func CheckAge(files []internal.FileInfo, rules config.Rules) (lowAge, mediumAge, highAge []internal.FileInfo) {
	lowAge = make([]internal.FileInfo, 0)
	mediumAge = make([]internal.FileInfo, 0)
	highAge = make([]internal.FileInfo, 0)

	for _, file := range files {
		if file.LastAccess < int16(rules.Age.Low) {
			lowAge = append(lowAge, file)
		} else if file.LastAccess < int16(rules.Age.Medium) {
			mediumAge = append(mediumAge, file)
		} else {
			highAge = append(highAge, file)
		}
	}

	return lowAge, mediumAge, highAge
}
