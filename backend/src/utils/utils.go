package utils

import (
	"log"
	"os"
)

func CreateFile(dirPath string, fileName string) (*os.File, error) {
	log.Printf("CreateFile by dir %s, with fileName %s", dirPath, fileName)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		log.Fatalf("os.MkdirAll() failed for dirPath %s, err %s",
			err, dirPath)
		return nil, err
	}

	allPath := dirPath + fileName
	file, err := os.Create(allPath)
	if err != nil {
		log.Fatalf("os.Create() failed for %s", allPath)
		return nil, err
	}

	return file, nil
}
