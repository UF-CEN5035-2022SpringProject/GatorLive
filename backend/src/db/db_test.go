package db

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
)

func pathSetup() {
	cwd, _ := os.Getwd()
	parent := filepath.Dir(cwd)
	os.Chdir(parent)
	fmt.Println("tt:" + cwd)
}

func TestSetUpEnv(t *testing.T) {
	pathSetup()
	logger.InitLogger()
}

// check the connection function
func TestDbConnection(t *testing.T) {
	ConnectionCreate()
}

func TestMapJWTObj(t *testing.T) {
	jwtToken := "test"
	jwtMap := MapJwtToken(jwtToken)
	if jwtMap == nil {
		t.Errorf("unable to get jwt obj from db")
	}
	logger.DebugLogger.Printf("[Test] TestJWTObj %v", jwtMap)
}
func TestGetUserObj(t *testing.T) {
	email := "test"
	userData := GetUserObj(email)
	if userData == nil {
		t.Errorf("unable to get user obj from db")
	}
	logger.DebugLogger.Printf("[Test] TestGetUserObj %v", userData)
}

func TestGetStoreObjbyUser(t *testing.T) {
	userId := "test"
	storeObj := GetStoreObjbyUserId(userId)
	if storeObj == nil {
		t.Errorf("unable to get store obj from db")
	}
	logger.DebugLogger.Printf("[Test] TestGetStoreObjbyUser %v", storeObj)
}

func TestGetStoreObj(t *testing.T) {
	storeId := "test"
	storeObj := GetStoreObj(storeId)
	if storeObj == nil {
		t.Errorf("unable to get jwt obj from db")
	}
	logger.DebugLogger.Printf("[Test] TestGetStoreObj %v", storeObj)
}
func TestLiveObj(t *testing.T) {
	liveId := "test"
	liveObj := GetLiveObj(liveId)
	if liveObj == nil {
		t.Errorf("unable to get live obj from db")
	}
	logger.DebugLogger.Printf("[Test] TestLiveObj %v", liveObj)
}
