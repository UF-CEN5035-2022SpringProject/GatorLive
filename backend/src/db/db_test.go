package db

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
	utils "github.com/UF-CEN5035-2022SpringProject/GatorStore/utils"
)

func pathSetup() {
	cwd, _ := os.Getwd()
	parent := filepath.Dir(cwd)
	os.Chdir(parent)
}

func TestSetUpEnv(t *testing.T) {
	pathSetup()
	logger.InitLogger()
}

// check the connection function
func TestDbConnection(t *testing.T) {
	ConnectionCreate()
}

func TestJWTObj(t *testing.T) {
	jwtToken := "test"
	jwtMap := MapJwtToken(jwtToken)
	if jwtMap == nil {
		t.Errorf("unable to get jwt obj from db")
	}
	logger.DebugLogger.Printf("[Test] TestJWTObj %v", jwtMap)
}

func TestUserObj(t *testing.T) {
	// Fully Test DB user functions Get, Add, Update, Delete
	email := "test"
	userData := GetUserObj(email)
	if userData != nil {
		t.Errorf("dirty test user data still in db")
	}

	// Test Add User Obj
	nowTime := time.Now().UTC().Format(time.RFC3339)
	testUserObj := UserObject{
		Id:          "test",
		Name:        "test",
		Email:       email,
		JwtToken:    utils.CreateJwtToken("test", "test", "test"),
		AccessToken: "testToken0",
		CreateTime:  nowTime,
		UpdateTime:  nowTime,
	}

	var saveMap map[string]interface{}
	userObjStr, _ := json.Marshal(testUserObj)
	json.Unmarshal(userObjStr, &saveMap)
	AddUserObj(email, saveMap)

	userData = GetUserObj(email)
	if userData == nil {
		t.Errorf("Add userObj Failed, unable to get user obj from db")
	}

	if !reflect.DeepEqual(saveMap, userData) {
		t.Errorf("Add userObj Failed, user obj retrieve not equal to save one")
	}

	updateToken := "testToken1"
	UpdateUserObj(email, "accessToken", updateToken)
	userData = GetUserObj(email)
	if userData == nil {
		t.Errorf("Update userObj Failed, unable to get user obj from db")
	}
	if userData["accessToken"] != updateToken {
		t.Errorf("Update userObj Failed, data did not update")
	}

	DeleteUserObj(email)
	userData = GetUserObj(email)
	if userData != nil {
		t.Errorf("Delete userObj Failed")
	}
}

func TestStoreObjbyUser(t *testing.T) {
	userId := "test"
	storeObj := GetStoreObjbyUserId(userId)
	if storeObj == nil {
		t.Errorf("unable to get store obj from db")
	}
	logger.DebugLogger.Printf("[Test] TestGetStoreObjbyUser %v", storeObj)
}

func TestStoreObj(t *testing.T) {
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
