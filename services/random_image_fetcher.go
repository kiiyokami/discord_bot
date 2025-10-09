package services

import (
	"math/rand"
	"os"
	"path/filepath"
)

func RandomImageFetcher(path string) (string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}
	if len(files) == 0 {
		allFiles := []string{}
		filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() {
				allFiles = append(allFiles, p)
			}
			return nil
		})
		if len(allFiles) == 0 {
			return "", nil
		}
		return allFiles[rand.Intn(len(allFiles))], nil
	}
	return files[rand.Intn(len(files))].Name(), nil
}
