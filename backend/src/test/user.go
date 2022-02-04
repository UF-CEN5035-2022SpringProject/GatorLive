package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	db "github.com/UF-CEN5035-2022SpringProject/GatorStore/db"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
)

func TestDBGetUserObj(response http.ResponseWriter, r *http.Request) {
	logger.DebugLogger.Printf("%v", db.FireBaseClient)

	dsnap, err := db.FireBaseClient.Collection(db.Collections["users"]).Doc("test").Get(db.DatabaseCtx)
	if err != nil {
		// t.Errorf("Error retreiving value in TestDBGetUserObj")
		logger.ErrorLogger.Fatalf("Error retreiving value in TestDBGetUserObj. %s", err)
	}
	m := dsnap.Data()
	logger.DebugLogger.Printf("Document data: %#v\n", m)
	jsonString, err := json.Marshal(m)
	fmt.Println(err)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonString)
}

func TestDBAddUserObj(*testing.T) {

}

func TestDBDeleteUserObj(*testing.T) {

}
