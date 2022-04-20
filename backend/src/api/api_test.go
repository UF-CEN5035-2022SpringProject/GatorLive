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

// TESTING MIDDLEWARE
// simple response
func EchoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GatorStore Backend is alive")
}
func TestHeaderMiddleWare(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/echo", EchoString)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/echo", nil)
	res, _ := http.DefaultClient.Do(req)

	if want, got := http.StatusOK, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	contentType := res.Header.Get("Content-Type")
	if want, got := "application/json", contentType; want != got {
		t.Fatalf("expected a %v, instead got: %v", want, got)
	}
}
func TestCrossAllowMiddleware(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/echo", EchoString)
	r.Use(CrossAllowMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodOptions, ts.URL+"/echo", nil)
	res, _ := http.DefaultClient.Do(req)

	if want, got := http.StatusOK, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}
func TestAuthMiddlewareNOAuth(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/echo", EchoString)
	r.Use(AuthMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/echo", nil)
	res, _ := http.DefaultClient.Do(req)
	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("Empty JWT")
	expectedUserData, _ := RespJSON{int(utils.MissingJwtTokenCode), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusUnauthorized, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}

func TestAuthMiddlewareNotExistJwtToken(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/echo", EchoString)
	r.Use(AuthMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/echo", nil)
	req.Header.Set("Authorization", "not exist")
	res, _ := http.DefaultClient.Do(req)
	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("Invalid JWT, not found")
	expectedUserData, _ := RespJSON{int(utils.MissingJwtTokenCode), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusUnauthorized, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}

func TestAuthMiddlewareValidJwtToken(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/echo", EchoString)
	r.Use(AuthMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/echo", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)

	if want, got := http.StatusOK, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestRRR(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/echo", EchoString)

	ts := httptest.NewServer(r)
	defer ts.Close()
	res, _ := http.Get(ts.URL + "/echo")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	greeting, _ := io.ReadAll(res.Body)
	res.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Printf("%s", greeting)
}

// TESTING USER

func TestGetUserProfile(t *testing.T) {
	// svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//     fmt.Fprintf(w, "%v", expect)
	// }))
	userTest := db.GetUserObj("test")
	// defer svr.Close()
	// c := NewClient(svr.URL)
	accessToken := fmt.Sprintf("%s", userTest["accessToken"])
	// if no accesstoken
	expect := &Profile{}
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

func TestUserOrderList(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/user/{userId}/order-list", UserOrderList)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/user/test/order-list", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)

	if want, got := http.StatusOK, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestUserOrderListNotAuthorized(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/user/{userId}/order-list", UserOrderList)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/user/notexist/order-list", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)
	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("invald request, permission denied")
	expectedUserData, _ := RespJSON{int(utils.InvalidJwtTokenCode), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusForbidden, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}

}

func TestUserOrderListInvalidPage(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/user/{userId}/order-list", UserOrderList)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/user/test/order-list?page=hi", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)
	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("Error type of page query")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusBadRequest, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}

}

func TestUserStoreList(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/user/{userId}/store-list", UserStoreList)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/user/test/store-list", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)

	if want, got := http.StatusOK, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestUserStoreListNotAuthorized(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/user/{userId}/store-list", UserStoreList)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/user/notexist/store-list", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)
	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("invald request, permission denied")
	expectedUserData, _ := RespJSON{int(utils.InvalidJwtTokenCode), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusForbidden, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}

}

func TestUserStoreListInvalidPage(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/user/{userId}/store-list", UserStoreList)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/user/test/store-list?page=hi", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)
	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("Error type of page query")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusBadRequest, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}

}

func TestUserInfo(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/user/{userId}/info", UserInfo)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/user/test/info", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)

	if want, got := http.StatusOK, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestUserInfoNotAuthorized(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/user/{userId}/info", UserInfo)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/user/notexsit/info", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)
	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("invald request, permission denied")
	expectedUserData, _ := RespJSON{int(utils.InvalidJwtTokenCode), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusForbidden, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}

}

func TestUserInfoPut(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/user/{userId}/info", UserInfo)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodPut, ts.URL+"/user/notexsit/info", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)
	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	expectStr := "Update notexsit user info"
	if want, got := http.StatusOK, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}

}

// TESTING LIVESTREAM

func TestCreateLivebroadcastNotAuthorized(t *testing.T) {
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

func TestCreateLivebroadcastInvalidStore(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/store/{storeId}/livestream", CreateLivebroadcast)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	data := &Title{
		Title: "123",
	}
	codeByte, _ := json.Marshal(data)
	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/store/noexist/livestream", strings.NewReader(string(codeByte)))
	req.Header.Add("Authorization", "test")

	res, _ := http.DefaultClient.Do(req)

	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("invald request, unable to get store")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()

	expectStr := string(expectedUserData)

	if want, got := http.StatusBadRequest, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}

	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}

func TestCreateLivebroadcastWrongStore(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/store/{storeId}/livestream", CreateLivebroadcast)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	data := &Title{
		Title: "123",
	}
	codeByte, _ := json.Marshal(data)
	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/store/gatorstore-1/livestream", strings.NewReader(string(codeByte)))
	req.Header.Add("Authorization", "test")

	res, _ := http.DefaultClient.Do(req)

	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("invald request, permission denied")
	expectedUserData, _ := RespJSON{int(utils.InvalidJwtTokenCode), errorMsg}.SetResponse()

	expectStr := string(expectedUserData)

	if want, got := http.StatusForbidden, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}

	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
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

func TestGetLiveStreamNotExistLiveId(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/live/status?liveId=notexist", nil)
	// req.Header.Add("Authorization", "test")
	w := httptest.NewRecorder()

	GetLiveStream(w, req)

	b, _ := io.ReadAll(w.Result().Body)

	resStr := string(b)

	errorMsg := utils.SetErrorMsg("invald request, unable to get livestream")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()

	expectStr := string(expectedUserData)

	if want, got := http.StatusBadRequest, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}

// TESTING PRODUCT

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
func TestProductUpdateNotAuthorized(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/product/{productId}/info", ProductRESTFUL)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodPut, ts.URL+"/product/product-1/info", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)
	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("jwtToken and storeId not match")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusForbidden, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}

func TestProductUpdateNotExistProductId(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/product/{productId}/info", ProductRESTFUL)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodPut, ts.URL+"/product/notexist/info", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)
	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("jwtToken and storeId not match")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusForbidden, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}

func TestProductUpdate(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/product/{productId}/info", ProductRESTFUL)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	p := 999.99
	q := 101
	data := &ProductUpdateObject{
		Description: "up",
		Name:        "upName",
		Picture:     "somePic",
		Price:       &p,
		Quantity:    &q,
	}
	codeByte, _ := json.Marshal(data)

	req, _ := http.NewRequest(http.MethodPut, ts.URL+"/product/product-11/info", strings.NewReader(string(codeByte)))
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)

	if want, got := http.StatusOK, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestProductUpdateEncodeFail(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/product/{productId}/info", ProductRESTFUL)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodPut, ts.URL+"/product/product-11/info", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)

	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("error occurs before ProductUpdate")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusInternalServerError, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}

func TestProductPurchase(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/product/{productId}/info", ProductRESTFUL)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	data := &ProductPurchaseReqObject{
		Quantity: 1,
		LiveId:   "",
	}
	codeByte, _ := json.Marshal(data)

	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/product/product-11/info", strings.NewReader(string(codeByte)))
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)

	if want, got := http.StatusOK, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestProductPurchaseTooMuch(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/product/{productId}/info", ProductRESTFUL)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	data := &ProductPurchaseReqObject{
		Quantity: 1000000,
		LiveId:   "",
	}
	codeByte, _ := json.Marshal(data)

	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/product/product-11/info", strings.NewReader(string(codeByte)))
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)

	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("Purchase too much")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusForbidden, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}
func TestProductPurchaseNoExistProduct(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/product/{productId}/info", ProductRESTFUL)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	data := &ProductPurchaseReqObject{
		Quantity: 1,
		LiveId:   "",
	}
	codeByte, _ := json.Marshal(data)

	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/product/notexist/info", strings.NewReader(string(codeByte)))
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)

	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("Unable to get product")
	expectedUserData, _ := RespJSON{int(utils.UnableToGetDbObj), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusNotFound, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}

func TestProductPurchaseBodyEncodeFail(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/product/{productId}/info", ProductRESTFUL)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/product/notexist/info", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)

	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("error occurs before ProductPurchase")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusInternalServerError, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}

func TestProductDelete(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/product/{productId}/info", ProductRESTFUL)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodDelete, ts.URL+"/product/product-12/info", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)

	if want, got := http.StatusOK, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestProductDeleteProductNotExist(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/product/{productId}/info", ProductRESTFUL)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodDelete, ts.URL+"/product/notexist/info", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)

	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("jwtToken and storeId not match")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusForbidden, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}

// TESTING STORE

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

func TestStoreOrder(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/store/{storeId}/order-list", StoreOrders)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/store/test/order-list", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)

	if want, got := http.StatusOK, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}

func TestStoreOrderInvalidStoreId(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/store/{storeId}/order-list", StoreOrders)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/store/notexist/order-list", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)

	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("invald request, unable to get store")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusBadRequest, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}

func TestStoreOrderWrongStoreId(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/store/{storeId}/order-list", StoreOrders)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/store/gatorstore-1/order-list", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)

	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("invald request, permission denied")
	expectedUserData, _ := RespJSON{int(utils.InvalidJwtTokenCode), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusForbidden, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}

func TestStoreOrderInvalidPageNum(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/store/{storeId}/order-list", StoreOrders)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/store/test/order-list?page=hi", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)

	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("Error type of page query")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusBadRequest, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}

func TestStoreUpdateIsLive(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/store/{storeId}/livestream/update", UpdateIsLive)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	data := &Status{
		IsLive: false,
	}
	ts := httptest.NewServer(r)
	defer ts.Close()
	codeByte, _ := json.Marshal(data)

	req, _ := http.NewRequest(http.MethodPut, ts.URL+"/store/test/livestream/update", strings.NewReader(string(codeByte)))
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)

	if want, got := http.StatusOK, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}
func TestStoreUpdateIsLiveNotExistStoreId(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/store/{storeId}/livestream/update", UpdateIsLive)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodPut, ts.URL+"/store/notexist/livestream/update", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)

	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("invald request, unable to get store")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusBadRequest, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}
func TestStoreUpdateIsLiveWrongStoreId(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/store/{storeId}/livestream/update", UpdateIsLive)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodPut, ts.URL+"/store/gatorstore-1/livestream/update", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)

	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("invald request, permission denied")
	expectedUserData, _ := RespJSON{int(utils.InvalidJwtTokenCode), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusForbidden, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}
func TestStoreUpdateIsLiveEncodeFail(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/store/{storeId}/livestream/update", UpdateIsLive)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthMiddleware)
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodPut, ts.URL+"/store/test/livestream/update", nil)
	req.Header.Set("Authorization", "test")
	res, _ := http.DefaultClient.Do(req)

	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("Unable to decode livestream status req")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusBadRequest, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}
func TestStoreRecommendList(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/store/recommend-list", StoreRecommendList)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/store/recommend-list", nil)
	res, _ := http.DefaultClient.Do(req)

	if want, got := http.StatusOK, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}
func TestStoreRecommendListInvalidPageNum(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/store/recommend-list", StoreRecommendList)
	r.Use(CrossAllowMiddleware)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(HeaderMiddleware)

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/store/recommend-list?page=hi", nil)
	res, _ := http.DefaultClient.Do(req)

	b, _ := io.ReadAll(res.Body)
	resStr := string(b)

	errorMsg := utils.SetErrorMsg("Error type of page query")
	expectedUserData, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
	expectStr := string(expectedUserData)
	if want, got := http.StatusBadRequest, res.StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	if resStr != expectStr {
		t.Errorf("expected res to be %v got %v", expectStr, resStr)
	}
}
