package test

import (
	"net/http"
	"testing"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/db"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
)

func TestDBGetUserObj(w http.ResponseWriter, r *http.Request) {
	dsnap, err := db.FireBaseClient.Collection(db.Collections["users"]).Doc("SF").Get(db.DatabaseCtx)
	if err != nil {
		// t.Errorf("Error retreiving value in TestDBGetUserObj")
		logger.ErrorLogger.Printf("Error retreiving value in TestDBGetUserObj")
	}
	m := dsnap.Data()
	logger.DebugLogger.Printf("Document data: %#v\n", m)
}

func TestDBAddUserObj(*testing.T) {

}

func TestDBDeleteUserObj(*testing.T) {

}
