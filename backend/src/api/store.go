package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/db"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/utils"
	gorillaContext "github.com/gorilla/context"
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
	IsLive     bool   `json:"isLive"`
	LiveId     string `json:"liveId"`
}

func StoreCreate(w http.ResponseWriter, r *http.Request) {
	// get code or assesstoken from http.request
	userData := gorillaContext.Get(r, "userData").(map[string]interface{})

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

	if rq.Name == "" {
		logger.ErrorLogger.Printf("Empty name to create store: %v", err)
		errorMsg := utils.SetErrorMsg("Empty name to create store")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}

	// create store object with name
	newStoreCount := db.GetStoreNewCount()
	newStoreId := "GatorStore_" + strconv.Itoa(newStoreCount)

	nowTime := time.Now().UTC().Format(time.RFC3339)
	storeObj := &StoreObject{
		Id:         newStoreId,
		Name:       rq.Name,
		UserId:     userData["id"].(string),
		CreateTime: nowTime,
		UpdateTime: nowTime,
		IsLive:     false,
		LiveId:     "",
	}

	var convertMap map[string]interface{}
	storeObjStr, _ := json.Marshal(storeObj)
	json.Unmarshal(storeObjStr, &convertMap)
	storeData := convertMap

	db.AddStoreObj(newStoreId, storeData)
	db.UpdateStoreCount(newStoreCount)

	resp, err := RespJSON{0, storeData}.SetResponse()
	if err != nil {
		logger.ErrorLogger.Printf("Error on wrapping JSON resp, err: %v", err)
		errorMsg := utils.SetErrorMsg("Error on wrapping JSON resp")
		resp, _ := RespJSON{int(utils.InvalidAccessTokenCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusInternalServerError)
		return
	}
	ReturnResponse(w, resp, http.StatusOK)
}
