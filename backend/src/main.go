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
	port    string = "8000"
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
	authApis.HandleFunc("/user/store-list", test.EchoString)

	// Store
	authApis.HandleFunc("/store/create", api.StoreCreate).Methods("POST", "OPTIONS")
	authApis.HandleFunc("/store/{storeId}/product-list", test.EchoString)
	authApis.HandleFunc("/store/{storeId}/livestream", api.CreateLivebroadcast).Methods("GET", "POST", "OPTIONS")
	authApis.HandleFunc("/store/{storeId}/livestreamStatus", api.LivestreamStatus).Methods("GET", "PUT", "OPTIONS")

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
	// testAuthApis.Use(api.AuthMiddleware)
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
