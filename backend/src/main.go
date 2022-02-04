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
)

func main() {
	logger.InitLogger()

	r := mux.NewRouter()
	// set up routing path
	prodRoutePrefix := "/api"

	// TEST API path
	testRoutePrefix := "/test/api"
	r.HandleFunc(testRoutePrefix+"/test", test.EchoString)
	r.HandleFunc(testRoutePrefix+"/test/searchUser", test.TestDBGetUserObj)

	// USER path
	r.HandleFunc(prodRoutePrefix+"/user/login", api.Login)
	r.HandleFunc(prodRoutePrefix+"/user/{userId}/info", api.UserInfo).Methods("GET", "PUT") // TODO missing authentication middleware
	r.HandleFunc(prodRoutePrefix+"/user/{userId}/store-list", test.EchoString)

	// Store
	r.HandleFunc(prodRoutePrefix+"/store/{storeId}/product-list", test.EchoString)

	// read google oauth2 credentials
	api.ReadCredential()
	logger.InfoLogger.Println("client id: " + api.ClientID)
	logger.InfoLogger.Println("client secret: " + api.ClientSecret)
	logger.InfoLogger.Println("redirect uris: " + strings.Join(api.RedirectURL, ","))

	// create DB connection
	db.ConnectionSetUp()
	//
	logger.InfoLogger.Println(appName + " server is start at port: " + port)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
