package main

import (
	"log"
	"net/http"
	"time"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/test"
)

var (
	port    string = "8080"
	appName string = "GatorStore"
)

func main() {
	logger.InitLogger()

	s := &http.Server{
		Addr:           ":" + port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// set up routing path
	http.HandleFunc("/test", test.EchoString)

	logger.InfoLogger.Println(appName + " server is start at port: " + port)
	log.Fatal(s.ListenAndServe())
}
