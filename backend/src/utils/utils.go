package utils

import (
	b64 "encoding/base64"
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

func CreateJwtToken(userId string, userEmail string, nowTime string) string {
	// store newJwt in DB
	newJwtToken := "gst." + b64.StdEncoding.EncodeToString([]byte(JwtPrefix+userEmail+userId)) + "_" + b64.StdEncoding.EncodeToString([]byte(nowTime))
	return newJwtToken
}

func Pagenator(targetSlice []map[string]interface{}, currectPage int, sliceSize int) []map[string]interface{} {
	low := currectPage * PageLimit
	high := (currectPage + 1) * PageLimit
	if high > sliceSize {
		high = sliceSize
	}
	return targetSlice[low:high]
}
