package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/db"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/utils"
	"github.com/gorilla/mux"
	// "fmt"
)

func pathSetup() {
	cwd, _ := os.Getwd()
	parent := filepath.Dir(cwd)
	os.Chdir(parent)
}

func TestSetUpEnv(t *testing.T) {
	pathSetup()
	logger.InitLogger()
	db.ConnectionCreate()
	ReadCredential()
}

func TestGetUserProfile(t *testing.T) {
	expect := &Profile{
		Name:  "YiMing Chang",
		Email: "yimingchang@ufl.edu",
	}
	// svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//     fmt.Fprintf(w, "%v", expect)
	// }))
	userTest := db.GetUserObj("test")
	// defer svr.Close()
	// c := NewClient(svr.URL)
	accessToken := fmt.Sprintf("%s", userTest["accessToken"])
	// if no accesstoken
	expect = &Profile{}
	res := GetUserProfile(accessToken)
	if res != *expect {
		t.Errorf("expected res to be %v got %v", *expect, res)
	}
}

func TestLogin(t *testing.T) {
	data := &Code{
		Code: "4/0AX4XfWhXHAUcU6v5oBWSGC5sxEwRkHdfjaEgGv4blsqJJxphuEtUVpp4ur7ZJNl-q8O7kw",
	}
	codeByte, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPost, "http://localhost/", strings.NewReader(string(codeByte)))
	w := httptest.NewRecorder()

	Login(w, req)

	errorMsg := utils.SetErrorMsg("Exchange token by code failed!")
	expectedUserData, _ := RespJSON{int(utils.InvalidGoogleCode), errorMsg}.SetResponse()

	b, _ := io.ReadAll(w.Result().Body)

	resStr := string(b)
	expectStr := string(expectedUserData)
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}

func TestCreateLivebroadcast(t *testing.T) {
	data := &Title{
		Title: "123",
	}
	codeByte, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPost, "/store/1/livestream", strings.NewReader(string(codeByte)))
	req.Header.Add("Authorization", "test")
	w := httptest.NewRecorder()
	vars := map[string]string{
		"storeId": "1",
	}
	req = mux.SetURLVars(req, vars)
	CreateLivebroadcast(w, req)

	if want, got := http.StatusUnauthorized, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestLivestreamStatus(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/store/1/livestreamStatus", nil)
	req.Header.Add("Authorization", "test")
	w := httptest.NewRecorder()
	vars := map[string]string{
		"storeId": "1",
	}
	req = mux.SetURLVars(req, vars)
	LivestreamStatus(w, req)

	expect := make(map[string]interface{})
	expect["createTime"] = "2006-01-02T15:04:05Z07:00"
	expect["id"] = "test"
	expect["isLive"] = false
	expect["name"] = "testStore"
	expect["updateTime"] = "2006-01-02T15:04:05Z07:00"
	expect["userID"] = "test"

	b, _ := RespJSON{0, expect}.SetResponse()
	expectStr := string(b)
	b, _ = io.ReadAll(w.Result().Body)
	resStr := string(b)
	logger.DebugLogger.Printf("expected a %v, resp: %v", expectStr, resStr)

	if expectStr != resStr {
		t.Fatalf("expected a %v, instead got: %v", expectStr, resStr)
	}
}
