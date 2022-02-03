package api

import (
	"fmt"
	"net/http"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
	"github.com/gorilla/mux"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// TODO @chouhy
	logger.DebugLogger.Println("User ___ Login")
}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	// Depend on the action
	// 1. Get userInfo
	logger.DebugLogger.Println(r.Method)
	vars := mux.Vars(r)
	if r.Method == "GET" {
		fmt.Fprintf(w, "Get %v user info", vars["userId"])
	} else if r.Method == "PUT" {
		fmt.Fprintf(w, "Update %v user info", vars["userId"])
	}
}
