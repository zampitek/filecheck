package checks

import "github.com/zampitek/filecheck/internal"

func CheckSize(files []internal.FileInfo) (lowSize, mediumSize, highSize []internal.FileInfo) {
	lowSize = make([]internal.FileInfo, 0)
	mediumSize = make([]internal.FileInfo, 0)
	highSize = make([]internal.FileInfo, 0)

	for _, file := range files {
		if file.Size < 100*1024*1024 {
			lowSize = append(lowSize, file)
		} else if file.Size < 1024*1024*1024 {
			mediumSize = append(mediumSize, file)
		} else {
			highSize = append(highSize, file)
		}
	}

	return lowSize, mediumSize, highSize
}
