package api

import (
	"fmt"
	"net/http"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/db"
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
		value := db.GetUserObj(vars["userId"])
		resp, err := JsonResponse(value, 0)
		if err != nil {
			logger.ErrorLogger.Fatalf("Error on wrapping JSON resp %s", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)

	} else if r.Method == "PUT" {
		fmt.Fprintf(w, "Update %v user info", vars["userId"])
	}
}
