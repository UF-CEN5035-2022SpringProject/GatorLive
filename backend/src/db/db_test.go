package db

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
)

func dbSetup() {
	cwd, _ := os.Getwd()
	parent := filepath.Dir(cwd)
	os.Chdir(parent)
	ConnectionCreate()
}

func loggerSetup() {
	cwd, _ := os.Getwd()
	parent := filepath.Dir(cwd)
	os.Chdir(parent)
	logger.InitLogger()
}

// check the connection function
func TestDbConnection(t *testing.T) {
	cwd, _ := os.Getwd()
	parent := filepath.Dir(cwd)
	os.Chdir(parent)
	ConnectionCreate()
}
func TestGetUserObj(t *testing.T) {
	dbSetup()
	loggerSetup()
	email := "test"
	userData := GetUserObj(email)
	if userData == nil {
		t.Errorf("unable to get user obj from db")
	}
	logger.DebugLogger.Printf("TestGetStoreObj %v", userData)
}
