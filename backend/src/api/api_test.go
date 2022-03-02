package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/db"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
	// "fmt"
)

func TestGetUserProfile(t *testing.T) {
	expect := &Profile{
		Name:  "Hung-You Chou",
		Email: "jimchou1995@gmail.com",
	}
	// svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//     fmt.Fprintf(w, "%v", expect)
	// }))

	// defer svr.Close()
	// c := NewClient(svr.URL)
	accessToken := ""
	res := GetUserProfile(accessToken)
	if res != *expect {
		t.Errorf("expected res to be %v got %v", *expect, res)
	}
}
func dbSetup() {
	os.Chdir("/home/chouhy/GatorStore/backend/src/")
	db.ConnectionCreate()
}
func loggerSetup() {
	os.Chdir("/home/chouhy/GatorStore/backend/src/")
	logger.InitLogger()
}
func apiSetup() {
	os.Chdir("/home/chouhy/GatorStore/backend/src/")
	ReadCredential()
}
func TestLogin(t *testing.T) {
	loggerSetup()
	dbSetup()
	apiSetup()
	data := &Code{
		Code: "4/0AX4XfWhXHAUcU6v5oBWSGC5sxEwRkHdfjaEgGv4blsqJJxphuEtUVpp4ur7ZJNl-q8O7kw",
	}
	codeByte, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPost, "http://localhost/", strings.NewReader(string(codeByte)))
	w := httptest.NewRecorder()

	Login(w, req)

	expectedUserData := make(map[string]interface{})
	expectedUserData["createTime"] = "2022-02-22T02:29:37Z"
	expectedUserData["email"] = "jimchou1995@gmail.com"
	expectedUserData["updateTime"] = "2022-02-22T02:29:37Z"
	expectedUserData["jwtToken"] = "gst.R2F0b3JTdG9yZV9qaW1jaG91MTk5NUBnbWFpbC5jb20xMTAwMw==_MjAyMi0wMi0yMlQwMjoyOTozN1o="
	expectedUserData["id"] = "11003"
	expectedUserData["name"] = "Hung-You Chou"

	b, _ := io.ReadAll(w.Result().Body)
	var resUserData map[string]interface{}
	// var temp db.UserObject
	// json.Unmarshal(b, &temp)
	// userObjStr, _ := json.Marshal(userObj)
	json.Unmarshal(b, &resUserData)
	delete(resUserData, "accessToken")

	resB, _ := json.Marshal(resUserData)
	expectB, _ := json.Marshal(expectedUserData)
	resStr := string(resB)
	expectStr := string(expectB)
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}
