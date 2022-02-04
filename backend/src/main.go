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

	r := mux.NewRouter()
	// set up routing path
	prodRoutePrefix := "/api"

	// TEST API path
	testRoutePrefix := "/test/api"
	r.HandleFunc(testRoutePrefix+"/test", test.EchoString).Methods("GET", "OPTIONS")
	r.HandleFunc(testRoutePrefix+"/user/login", test.TestDBGetUserObj).Methods("GET", "POST", "OPTIONS")

	// USER path
	r.HandleFunc(prodRoutePrefix+"/user/login", api.Login)
	r.HandleFunc(prodRoutePrefix+"/user/info", api.UserInfo).Methods("GET", "PUT", "OPTIONS") // TODO missing authentication middleware
	r.HandleFunc(prodRoutePrefix+"/user/store-list", test.EchoString)

	// Store
	r.HandleFunc(prodRoutePrefix+"/store/{storeId}/product-list", test.EchoString)

	// read google oauth2 credentials
	api.ReadCredential()
	logger.InfoLogger.Println("client id: " + api.ClientID)
	logger.InfoLogger.Println("client secret: " + api.ClientSecret)
	logger.InfoLogger.Println("redirect uris: " + strings.Join(api.RedirectURL, ","))

	// If debug = Ture then set the CORSMethodMiddleware
	if IsDev {
		r.Use(api.CrossAllowMiddleware)
		r.Use(mux.CORSMethodMiddleware(r))
	}
	r.Use(api.HeaderMiddleware)

	//
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
