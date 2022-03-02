package test

import (
	"net/http"
	"testing"

	api "github.com/UF-CEN5035-2022SpringProject/GatorStore/api"
	db "github.com/UF-CEN5035-2022SpringProject/GatorStore/db"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/utils"
)

func TestDBGetUserObj(w http.ResponseWriter, r *http.Request) {
	dsnap, err := db.FireBaseClient.Collection(db.DbCollections["users"]).Doc("test").Get(db.DatabaseCtx)
	if err != nil {
		// t.Errorf("Error retreiving value in TestDBGetUserObj")
		logger.ErrorLogger.Printf("Error retreiving value in TestDBGetUserObj. %s", err)
	}
	value := dsnap.Data()
	logger.DebugLogger.Printf("Document data: %#v\n", value)

	resp, err := api.RespJSON{0, value}.SetResponse()
	if err != nil {
		logger.ErrorLogger.Printf("Error on wrapping JSON resp %s", err)
		errorMsg := utils.SetErrorMsg("Error on wrapping JSON")
		resp, _ := api.RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		api.ReturnResponse(w, resp, http.StatusInternalServerError)
		return
	}
	api.ReturnResponse(w, resp, http.StatusOK)
}

func TestDBAddUserObj(*testing.T) {

}

func TestDBDeleteUserObj(*testing.T) {

}
