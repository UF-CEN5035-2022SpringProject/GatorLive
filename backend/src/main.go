package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/api"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/db"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/test"
	"github.com/gorilla/mux"
)

var (
	port    string = "8080"
	appName string = "GatorStore"
	IsDev   bool   = true
)

func main() {
	logger.InitLogger()
	// create DB connection
	db.ConnectionCreate()

	// set up root routing path
	prodRoutePrefix := "/api"
	testRoutePrefix := "/test/api"

	r := mux.NewRouter()

	// login API
	r.HandleFunc(prodRoutePrefix+"/user/login", api.Login).Methods("GET", "POST", "OPTIONS")
	authApis := r.PathPrefix(prodRoutePrefix).Subrouter()
	// USER path
	authApis.HandleFunc("/user/{userId}/info", api.UserInfo).Methods("GET", "PUT", "OPTIONS")
	authApis.HandleFunc("/user/{userId}/store-list", api.UserStoreList).Methods("GET", "OPTIONS")
	authApis.HandleFunc("/user/{userId}/order-list", api.UserOrderList).Methods("GET", "OPTIONS")

	// Store
	authApis.HandleFunc("/store/create", api.StoreCreate).Methods("POST", "OPTIONS")
	r.HandleFunc(prodRoutePrefix+"/store/{storeId}/info", api.StoreInfo).Methods("GET")
	r.HandleFunc(prodRoutePrefix+"/store/{storeId}/product-list", api.StoreProducts).Methods("GET", "OPTIONS")
	authApis.HandleFunc("/store/{storeId}/order-list", api.StoreOrders).Methods("GET", "OPTIONS")
	// authApis.HandleFunc("/store/{storeId}/live-list", api.).Methods("GET", "OPTIONS")

	authApis.HandleFunc("/store/{storeId}/livestream", api.CreateLivebroadcast).Methods("GET", "POST", "OPTIONS")
	r.HandleFunc(prodRoutePrefix+"/store/{storeId}/livestream/info", api.LiveStreamInfo).Methods("GET", "OPTIONS")
	authApis.HandleFunc("/store/{storeId}/livestream/update", api.UpdateIsLive).Methods("PUT", "OPTIONS")

	// Product
	authApis.HandleFunc("/product/create", api.ProductCreate).Methods("POST", "OPTIONS")
	authApis.HandleFunc("/product/{productId}", api.ProductRESTFUL).Methods("POST", "GET", "PUT", "DELETE", "OPTIONS")
	// TEST API path
	r.HandleFunc(testRoutePrefix+"/echo", test.EchoString).Methods("GET", "OPTIONS")
	r.HandleFunc(testRoutePrefix+"/user/info", test.TestDBGetUserObj).Methods("GET", "OPTIONS")
	// testAuthApis := r.PathPrefix(testRoutePrefix).Subrouter()
	//testAuthApis.HandleFunc("/user/info", test.TestDBGetUserObj)

	// read google oauth2 credentials
	api.ReadCredential()
	logger.InfoLogger.Println("client id: " + api.ClientID)
	logger.InfoLogger.Println("client secret: " + api.ClientSecret)
	logger.InfoLogger.Println("redirect uris: " + strings.Join(api.RedirectURL, ","))

	// If debug = True then set the CORSMethodMiddleware
	if IsDev {
		// r.Use(api.LoggingMiddleware)
		r.Use(api.CrossAllowMiddleware)
		r.Use(mux.CORSMethodMiddleware(r))
	}
	authApis.Use(api.AuthMiddleware)
	r.Use(api.HeaderMiddleware)

	logger.InfoLogger.Println(appName + " server is start at port: " + port)
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
