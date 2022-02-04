package utils

import (
	"log"
	"net/http"
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

type returnObj struct {
	status int    `json:"status"`
	result []byte `json:"result"`
}

func ReturnWrapper(response http.ResponseWriter, resultBody []byte, errorCode int) {
	// Return json object
	// jsonObj := returnObj{
	// 	status: errorCode,
	// 	result: resultBody,
	// }

	// response.Header().Set("Content-Type", "application/json")
	// response.WriteHeader(http.StatusOK)
	// response.Write(jsonString)
}
