package scanner

import (
	"os"
	"path/filepath"
	"time"
)

type FileInfo struct {
	Path       string
	Size       int64
	IsDir      bool
	LastAccess int16
}

type ScanReport struct {
	Files []FileInfo
}

func Scan(root string) (ScanReport, error) {
	var files []FileInfo

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		files = append(files, FileInfo{
			Path:       path,
			Size:       info.Size(),
			IsDir:      info.IsDir(),
			LastAccess: int16(time.Since(info.ModTime()).Hours() / 24),
		})

		return nil
	})

	return ScanReport{Files: files}, err
}
