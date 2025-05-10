package checks

import (
	"github.com/zampitek/filecheck/internal"
)

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
