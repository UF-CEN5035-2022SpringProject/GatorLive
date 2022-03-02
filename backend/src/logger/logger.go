package logger

import (
	"log"
	"os"
	"time"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/utils"
)

// Other package can export these logger
var (
	DebugLogger   *log.Logger
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func InitLogger() {
	cwd, _ := os.Getwd()
	dirPath := cwd + "/log/"
	fileName := "GatorStore-backend-" + time.Now().Format("2006-01-02") + ".log"

	file, err := os.OpenFile(dirPath+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		file, _ = utils.CreateFile(dirPath, fileName)
	}

	DebugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
