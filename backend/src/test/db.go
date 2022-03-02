package test

import (
	"net/http"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/api"
	db "github.com/UF-CEN5035-2022SpringProject/GatorStore/db"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/utils"
)

func TestDBGetUserObj(w http.ResponseWriter, r *http.Request) {
	if db.FireBaseClient == nil || db.DatabaseCtx == nil {
		db.ConnectionCreate(true)
	}

	dsnap, err := db.FireBaseClient.Collection(db.DbCollections["users"]).Doc("test").Get(db.DatabaseCtx)
	if err != nil {
		errorMsg := utils.SetErrorMsg("Test failed. Error retreiving user obj.")
		resp, _ := api.RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		api.ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}
	value := dsnap.Data()
	resp, err := api.RespJSON{0, value}.SetResponse()
	if err != nil {
		errorMsg := utils.SetErrorMsg("Test failed. Error retreiving user obj.")
		resp, _ := api.RespJSON{int(utils.UnknownInternalErrCode), errorMsg}.SetResponse()
		api.ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}
	api.ReturnResponse(w, resp, http.StatusOK)
}
