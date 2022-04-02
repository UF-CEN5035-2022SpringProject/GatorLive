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
	gorillaContext "github.com/gorilla/context"
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
	req := httptest.NewRequest(http.MethodPost, "/store/test/livestream", strings.NewReader(string(codeByte)))
	req.Header.Add("Authorization", "test")
	w := httptest.NewRecorder()
	vars := map[string]string{
		"storeId": "test",
	}
	req = mux.SetURLVars(req, vars)
	jwtMap := db.MapJwtToken("test")
	gorillaContext.Set(req, "userData", db.GetUserObj(jwtMap["Email"].(string)))
	CreateLivebroadcast(w, req)

	if want, got := http.StatusUnauthorized, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestGetLiveStream(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/live/status?detail=false&liveId=test", nil)
	// req.Header.Add("Authorization", "test")
	w := httptest.NewRecorder()

	GetLiveStream(w, req)

	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestProductCreate(t *testing.T) {
	data := &ProductCreateObject{
		Name:        "test",
		Price:       99.99,
		Description: "testing",
		Quantity:    25,
		Picture:     "111",
		StoreId:     "test",
	}
	codeByte, _ := json.Marshal(data)

	req := httptest.NewRequest(http.MethodPost, "/product/create", strings.NewReader(string(codeByte)))
	req.Header.Add("Authorization", "test")
	w := httptest.NewRecorder()
	jwtMap := db.MapJwtToken("test")
	gorillaContext.Set(req, "userData", db.GetUserObj(jwtMap["Email"].(string)))
	ProductCreate(w, req)

	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}

}
func TestProductCreateWrongMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/product/create", nil)
	w := httptest.NewRecorder()
	ProductCreate(w, req)
	errorMsg := utils.SetErrorMsg("wrong method")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()

	b, _ := io.ReadAll(w.Result().Body)

	resStr := string(b)
	expectStr := string(expectedUserData)
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}

func TestProductNoData(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, "/product/create", nil)
	req.Header.Add("Authorization", "test")
	w := httptest.NewRecorder()
	jwtMap := db.MapJwtToken("test")
	gorillaContext.Set(req, "userData", db.GetUserObj(jwtMap["Email"].(string)))
	ProductCreate(w, req)

	errorMsg := utils.SetErrorMsg("error occurs before ProductCreate")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
	b, _ := io.ReadAll(w.Result().Body)

	resStr := string(b)
	expectStr := string(expectedUserData)
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}
func TestProductCreateWrongStoreId(t *testing.T) {
	data := &ProductCreateObject{
		Name:        "test",
		Price:       99.99,
		Description: "testing",
		Quantity:    25,
		Picture:     "111",
		StoreId:     "wrongStoreId",
	}
	codeByte, _ := json.Marshal(data)

	req := httptest.NewRequest(http.MethodPost, "/product/create", strings.NewReader(string(codeByte)))
	req.Header.Add("Authorization", "test")
	w := httptest.NewRecorder()
	jwtMap := db.MapJwtToken("test")
	gorillaContext.Set(req, "userData", db.GetUserObj(jwtMap["Email"].(string)))
	ProductCreate(w, req)

	errorMsg := utils.SetErrorMsg("jwtToken and storeId not match")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
	b, _ := io.ReadAll(w.Result().Body)
	resStr := string(b)
	expectStr := string(expectedUserData)
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}
func TestProductGet(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/product/productId/info", nil)
	req.Header.Add("Authorization", "test")
	w := httptest.NewRecorder()
	vars := map[string]string{
		"productId": "test",
	}
	req = mux.SetURLVars(req, vars)
	ProductGet(w, req)

	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}
func TestProductGetNotExistProduct(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/product/productId/info", nil)
	req.Header.Add("Authorization", "test")
	w := httptest.NewRecorder()
	vars := map[string]string{
		"productId": "not exist",
	}
	req = mux.SetURLVars(req, vars)
	ProductGet(w, req)

	errorMsg := utils.SetErrorMsg("Unable to get product")
	expectedUserData, _ := RespJSON{int(utils.UnableToGetDbObj), errorMsg}.SetResponse()
	b, _ := io.ReadAll(w.Result().Body)
	resStr := string(b)
	expectStr := string(expectedUserData)
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
	if want, got := http.StatusNotFound, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}
func TestStoreCreate(t *testing.T) {
	data := &storeCreateBody{
		Name: "test123",
	}
	codeByte, _ := json.Marshal(data)

	req := httptest.NewRequest(http.MethodPost, "/store/create", strings.NewReader(string(codeByte)))
	req.Header.Add("Authorization", "test")
	w := httptest.NewRecorder()
	jwtMap := db.MapJwtToken("test")
	gorillaContext.Set(req, "userData", db.GetUserObj(jwtMap["Email"].(string)))
	StoreCreate(w, req)

	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestStoreCreateNoData(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/store/create", nil)
	req.Header.Add("Authorization", "test")
	w := httptest.NewRecorder()
	jwtMap := db.MapJwtToken("test")
	gorillaContext.Set(req, "userData", db.GetUserObj(jwtMap["Email"].(string)))
	StoreCreate(w, req)

	errorMsg := utils.SetErrorMsg("error occurs before store created")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
	b, _ := io.ReadAll(w.Result().Body)
	resStr := string(b)
	expectStr := string(expectedUserData)
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
	if want, got := http.StatusInternalServerError, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestStoreCreateEmptyName(t *testing.T) {
	data := &storeCreateBody{
		Name: "",
	}
	codeByte, _ := json.Marshal(data)

	req := httptest.NewRequest(http.MethodPost, "/store/create", strings.NewReader(string(codeByte)))
	req.Header.Add("Authorization", "test")
	w := httptest.NewRecorder()
	jwtMap := db.MapJwtToken("test")
	gorillaContext.Set(req, "userData", db.GetUserObj(jwtMap["Email"].(string)))
	StoreCreate(w, req)

	errorMsg := utils.SetErrorMsg("Empty name to create store")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
	b, _ := io.ReadAll(w.Result().Body)
	resStr := string(b)
	expectStr := string(expectedUserData)
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
	if want, got := http.StatusBadRequest, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestStoreInfo(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/store/storeId/info", nil)
	w := httptest.NewRecorder()
	vars := map[string]string{
		"storeId": "test",
	}
	req = mux.SetURLVars(req, vars)
	StoreInfo(w, req)

	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}
func TestStoreInfoWrongStoreId(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/store/storeId/info", nil)
	w := httptest.NewRecorder()
	vars := map[string]string{
		"storeId": "not found",
	}
	req = mux.SetURLVars(req, vars)
	StoreInfo(w, req)

	errorMsg := utils.SetErrorMsg("invald request, empty store")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
	b, _ := io.ReadAll(w.Result().Body)
	resStr := string(b)
	expectStr := string(expectedUserData)
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
	if want, got := http.StatusBadRequest, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestStoreProducts(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/store/storeId/product-list?page=10000", nil)
	w := httptest.NewRecorder()
	vars := map[string]string{
		"storeId": "test",
	}
	req = mux.SetURLVars(req, vars)
	StoreProducts(w, req)

	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}
func TestStoreProductsWrongStoreId(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/store/storeId/product-list?page=0", nil)
	w := httptest.NewRecorder()
	vars := map[string]string{
		"storeId": "not found",
	}
	req = mux.SetURLVars(req, vars)
	StoreProducts(w, req)

	errorMsg := utils.SetErrorMsg("invald request, unable to get store")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
	b, _ := io.ReadAll(w.Result().Body)
	resStr := string(b)
	expectStr := string(expectedUserData)
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
	if want, got := http.StatusBadRequest, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}
func TestStoreProductsPageNotNumber(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/store/storeId/product-list?page=a", nil)
	w := httptest.NewRecorder()
	vars := map[string]string{
		"storeId": "test",
	}
	req = mux.SetURLVars(req, vars)
	StoreProducts(w, req)

	errorMsg := utils.SetErrorMsg("Error type of page query")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
	b, _ := io.ReadAll(w.Result().Body)
	resStr := string(b)
	expectStr := string(expectedUserData)
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
	if want, got := http.StatusBadRequest, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}
