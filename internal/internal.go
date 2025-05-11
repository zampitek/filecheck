package internal

import (
	"os"
	"path/filepath"
	"sort"
	"time"
)

type FileInfo struct {
	Name       string
	Path       string
	Size       int64
	IsDir      bool
	LastAccess int16
	NumFiles   int64
}

func Scan(root string) ([]FileInfo, error) {
	var files []FileInfo

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		files = append(files, FileInfo{
			Name:       info.Name(),
			Path:       path,
			Size:       info.Size(),
			IsDir:      info.IsDir(),
			LastAccess: int16(time.Since(info.ModTime()).Hours() / 24),
		})

		return nil
	})

	return files, err
}

func SortByAge(files []FileInfo) []FileInfo {
	sort.Slice(files, func(i, j int) bool {
		return files[i].LastAccess > files[j].LastAccess
	})

	return files
}

func SortBySize(files []FileInfo) []FileInfo {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Size > files[j].Size
	})

	return files
}
