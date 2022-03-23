package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/db"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/utils"
)

type storeCreateBody struct {
	Name string `json:"name"`
}

type StoreObject struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	UserId     string `json:"userId"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	IsLive     string `json:"isLive"`
	LiveId     string `json:"liveId"`
}

func StoreCreate(w http.ResponseWriter, r *http.Request) {
	// get code or assesstoken from http.request
	b, err := io.ReadAll(r.Body)
	logger.DebugLogger.Printf("StoreCreate request body %s", b)

	if err != nil {
		logger.ErrorLogger.Printf("Unable to read StoreCreate request body, err: %v", err)
		errorMsg := utils.SetErrorMsg("error occurs before store created")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusInternalServerError)
		return
	}

	var rq storeCreateBody
	err = json.Unmarshal(b, &rq)
	if err != nil {
		logger.ErrorLogger.Printf("Unable to decode StoreCreate request body, err: %v", err)
		errorMsg := utils.SetErrorMsg("error occurs before store created")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusInternalServerError)
		return
	}

	// create store object with name
	newStoreCount := db.GetStoreNewCount()
	// newStoreId := "GatorStore_" + strconv.Itoa(newStoreCount)

	// storeData := &db.StoreObject{
	// 	Id:   newStoreId,
	// 	Name: rq.Name,
	// 	// UserId:
	// 	// CreateTime:  nowTime,
	// 	// UpdateTime:  nowTime,
	// }

	// db.AddStoreObj(newStoreId, "storeData")
	db.UpdateUserCount(newStoreCount)
}
