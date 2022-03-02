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
	// "fmt"
)

func TestGetUserProfile(t *testing.T) {
	loggerSetup()
	dbSetup()
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

func dbSetup() {
	cwd, _ := os.Getwd()
	parent := filepath.Dir(cwd)
	os.Chdir(parent)
	db.ConnectionCreate()
}

func loggerSetup() {
	cwd, _ := os.Getwd()
	parent := filepath.Dir(cwd)
	os.Chdir(parent)
	logger.InitLogger()
}

func apiSetup() {
	cwd, _ := os.Getwd()
	parent := filepath.Dir(cwd)
	os.Chdir(parent)
	ReadCredential()
}
func TestLogin(t *testing.T) {
	loggerSetup()
	// dbSetup()
	apiSetup()
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
	loggerSetup()
	dbSetup()
	data := &Title{
		Title: "123",
	}
	codeByte, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPost, "/store/1/livestream", strings.NewReader(string(codeByte)))
	req.Header.Add("Authorization", "test(Do not delete)")
	w := httptest.NewRecorder()

	CreateLivebroadcast(w, req)
}
