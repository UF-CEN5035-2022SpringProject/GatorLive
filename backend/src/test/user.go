package test

import (
	"net/http"
	"testing"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/api"
	db "github.com/UF-CEN5035-2022SpringProject/GatorStore/db"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
)

func TestDBGetUserObj(w http.ResponseWriter, r *http.Request) {
	dsnap, err := db.FireBaseClient.Collection(db.DbCollections["users"]).Doc("test").Get(db.DatabaseCtx)
	if err != nil {
		// t.Errorf("Error retreiving value in TestDBGetUserObj")
		logger.ErrorLogger.Fatalf("Error retreiving value in TestDBGetUserObj. %s", err)
	}
	value := dsnap.Data()
	logger.DebugLogger.Printf("Document data: %#v\n", value)

	resp, err := api.JsonResponse(value, 0)
	if err != nil {
		logger.ErrorLogger.Fatalf("Error on wrapping JSON resp %s", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func TestDBAddUserObj(*testing.T) {

}

func TestDBDeleteUserObj(*testing.T) {

}
