package services

import (
	"math/rand"
	"os"
)

func RandomImageFetcher(path string) (string, error) {
	files, err := os.ReadDir(path)
	if err != nil || len(files) == 0 {
		return "", err
	}
	return files[rand.Intn(len(files))].Name(), nil
}
