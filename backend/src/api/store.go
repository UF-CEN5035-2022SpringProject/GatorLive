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
	"github.com/gorilla/mux"
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
type Status struct {
	IsLive bool `json:"isLive"`
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
	newStoreId := "gatorstore-" + strconv.Itoa(newStoreCount)

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

func StoreInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if r.Method == "GET" {
		storeId := vars["storeId"]
		storeObj := db.GetStoreObj(storeId)

		if storeObj == nil {
			logger.ErrorLogger.Printf("invald request, empty store")
			errorMsg := utils.SetErrorMsg("invald request, empty store")
			resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
			ReturnResponse(w, resp, http.StatusBadRequest)
			return
		}

		resp, err := RespJSON{0, storeObj}.SetResponse()
		if err != nil {
			logger.ErrorLogger.Printf("Error on wrapping JSON resp, err: %v", err)
			errorMsg := utils.SetErrorMsg("Error on wrapping JSON resp")
			resp, _ := RespJSON{int(utils.InvalidAccessTokenCode), errorMsg}.SetResponse()
			ReturnResponse(w, resp, http.StatusInternalServerError)
			return
		}

		ReturnResponse(w, resp, http.StatusOK)
		return
	}
}

func StoreProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storeId := vars["storeId"]

	// check if request is valid to access store product list
	storeObj := db.GetStoreObj(storeId)
	if storeObj == nil {
		logger.ErrorLogger.Printf("invald request, unable to get store")
		errorMsg := utils.SetErrorMsg("invald request, unable to get store")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}

	page := r.URL.Query().Get("page")
	if page == "" {
		page = "0"
	}

	intPage, err := strconv.Atoi(page)
	if err != nil {
		logger.ErrorLogger.Printf("Error page type, err: %v", err)
		errorMsg := utils.SetErrorMsg("Error type of page query")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}

	productList := db.GetStoreProducts(storeId, intPage)

	storeProductData := make(map[string]interface{})
	storeProductData["storeId"] = storeId

	productListSize := len(productList)
	storeProductData["maxPage"] = 0
	storeProductData["currectPage"] = 0
	storeProductData["productList"] = productList

	logger.DebugLogger.Printf("storeListSize: %d", productListSize)
	if productListSize != 0 {
		totalPage := (productListSize / utils.PageLimit)
		if (productListSize % utils.PageLimit) != 0 {
			totalPage += 1
		}
		maxPage := totalPage - 1
		storeProductData["maxPage"] = maxPage

		currectPage := intPage
		if currectPage > maxPage {
			currectPage = maxPage
		}
		storeProductData["currectPage"] = currectPage
		// arrange the pagenate
		storeProductData["productList"] = utils.Pagenator(productList, currectPage, productListSize)
	}

	resp, err := RespJSON{0, storeProductData}.SetResponse()
	if err != nil {
		logger.ErrorLogger.Printf("Error on wrapping JSON resp, err: %v", err)
		errorMsg := utils.SetErrorMsg("Error on wrapping JSON resp")
		resp, _ := RespJSON{int(utils.InvalidAccessTokenCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusInternalServerError)
		return
	}

	ReturnResponse(w, resp, http.StatusOK)
}

func StoreOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storeId := vars["storeId"]

	// check if request is valid to access store product list
	storeObj := db.GetStoreObj(storeId)
	if storeObj == nil {
		logger.ErrorLogger.Printf("invald request, unable to get store")
		errorMsg := utils.SetErrorMsg("invald request, unable to get store")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}

	userData := gorillaContext.Get(r, "userData").(map[string]interface{})
	if storeObj["userId"] != userData["id"].(string) {
		logger.ErrorLogger.Printf("invald request, permission denied")
		errorMsg := utils.SetErrorMsg("invald request, permission denied")
		resp, _ := RespJSON{int(utils.InvalidJwtTokenCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusForbidden)
		return
	}

	page := r.URL.Query().Get("page")
	if page == "" {
		page = "0"
	}

	intPage, err := strconv.Atoi(page)
	if err != nil {
		logger.ErrorLogger.Printf("Error page type, err: %v", err)
		errorMsg := utils.SetErrorMsg("Error type of page query")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}

	orderList := db.GetStoreOrders(storeId, intPage)

	storeOrderData := make(map[string]interface{})
	storeOrderData["storeId"] = storeId

	orderListSize := len(orderList)
	storeOrderData["maxPage"] = 0
	storeOrderData["currectPage"] = 0
	storeOrderData["orderList"] = orderList

	if orderListSize != 0 {
		totalPage := (orderListSize / utils.PageLimit)
		if (orderListSize % utils.PageLimit) != 0 {
			totalPage += 1
		}
		maxPage := totalPage - 1
		storeOrderData["maxPage"] = maxPage

		currectPage := intPage
		if currectPage > maxPage {
			currectPage = maxPage
		}
		storeOrderData["currectPage"] = currectPage
		// arrange the pagenate
		storeOrderData["orderList"] = utils.Pagenator(orderList, currectPage, orderListSize)
	}

	resp, err := RespJSON{0, storeOrderData}.SetResponse()
	if err != nil {
		logger.ErrorLogger.Printf("Error on wrapping JSON resp, err: %v", err)
		errorMsg := utils.SetErrorMsg("Error on wrapping JSON resp")
		resp, _ := RespJSON{int(utils.InvalidAccessTokenCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusInternalServerError)
		return
	}

	ReturnResponse(w, resp, http.StatusOK)
}

func UpdateIsLive(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storeId := vars["storeId"]

	userData := gorillaContext.Get(r, "userData").(map[string]interface{})

	storeObj := db.GetStoreObj(storeId)
	if storeObj == nil {
		logger.ErrorLogger.Printf("invald request, unable to get store")
		errorMsg := utils.SetErrorMsg("invald request, unable to get store")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}

	userId := userData["id"].(string)
	if storeObj["userId"] != userId {
		logger.ErrorLogger.Printf("invald request, permission denied")
		errorMsg := utils.SetErrorMsg("invald request, permission denied")
		resp, _ := RespJSON{int(utils.InvalidJwtTokenCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusForbidden)
		return
	}

	// update store object
	b, err := io.ReadAll(r.Body)
	if err != nil {
		logger.ErrorLogger.Printf("Unable to read livestream status req: %v", err)
		errorMsg := utils.SetErrorMsg("Unable to read livestream status req")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}
	var status Status
	err = json.Unmarshal(b, &status)
	if err != nil {
		logger.ErrorLogger.Printf("Unable to decode livestream status req: %v", err)
		errorMsg := utils.SetErrorMsg("Unable to decode livestream status req")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}
	db.UpdateStoreObj(storeId, "isLive", status.IsLive)

	// get updated store object
	storeObj = db.GetStoreObj(storeId)
	resp, err := RespJSON{0, storeObj}.SetResponse()
	if err != nil {
		logger.ErrorLogger.Printf("Error on wrapping JSON resp, Error: %s", err)
	}
	ReturnResponse(w, resp, http.StatusOK)
	return
}

func StoreRecommendList(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("page")
	if page == "" {
		page = "0"
	}

	intPage, err := strconv.Atoi(page)
	if err != nil {
		logger.ErrorLogger.Printf("Error page type, err: %v", err)
		errorMsg := utils.SetErrorMsg("Error type of page query")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}

	storeList := db.GetStoreRecommend(intPage)

	storeRecommendData := make(map[string]interface{})

	storeListSize := len(storeList)
	storeRecommendData["maxPage"] = 0
	storeRecommendData["currectPage"] = 0
	storeRecommendData["storeList"] = storeList

	if storeListSize != 0 {
		totalPage := (storeListSize / utils.PageLimit)
		if (storeListSize % utils.PageLimit) != 0 {
			totalPage += 1
		}
		maxPage := totalPage - 1
		storeRecommendData["maxPage"] = maxPage

		currectPage := intPage
		if currectPage > maxPage {
			currectPage = maxPage
		}
		storeRecommendData["currectPage"] = currectPage
		// arrange the pagenate
		storeRecommendData["storeList"] = utils.Pagenator(storeList, currectPage, storeListSize)
	}

	resp, err := RespJSON{0, storeRecommendData}.SetResponse()
	if err != nil {
		logger.ErrorLogger.Printf("Error on wrapping JSON resp, err: %v", err)
		errorMsg := utils.SetErrorMsg("Error on wrapping JSON resp")
		resp, _ := RespJSON{int(utils.InvalidAccessTokenCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusInternalServerError)
		return
	}

	ReturnResponse(w, resp, http.StatusOK)
}

func StoreLives(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storeId := vars["storeId"]

	// check if request is valid to access store product list
	storeObj := db.GetStoreObj(storeId)
	if storeObj == nil {
		logger.ErrorLogger.Printf("invald request, unable to get store")
		errorMsg := utils.SetErrorMsg("invald request, unable to get store")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}

	userData := gorillaContext.Get(r, "userData").(map[string]interface{})
	if storeObj["userId"] != userData["id"].(string) {
		logger.ErrorLogger.Printf("invald request, permission denied")
		errorMsg := utils.SetErrorMsg("invald request, permission denied")
		resp, _ := RespJSON{int(utils.InvalidJwtTokenCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusForbidden)
		return
	}

	page := r.URL.Query().Get("page")
	if page == "" {
		page = "0"
	}

	intPage, err := strconv.Atoi(page)
	if err != nil {
		logger.ErrorLogger.Printf("Error page type, err: %v", err)
		errorMsg := utils.SetErrorMsg("Error type of page query")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}

	liveList := db.GetStoreLives(storeId, intPage)

	storeLiveData := make(map[string]interface{})
	storeLiveData["storeId"] = storeId

	liveListSize := len(liveList)
	storeLiveData["maxPage"] = 0
	storeLiveData["currectPage"] = 0
	storeLiveData["liveList"] = liveList

	if liveListSize != 0 {
		totalPage := (liveListSize / utils.PageLimit)
		if (liveListSize % utils.PageLimit) != 0 {
			totalPage += 1
		}
		maxPage := totalPage - 1
		storeLiveData["maxPage"] = maxPage

		currectPage := intPage
		if currectPage > maxPage {
			currectPage = maxPage
		}
		storeLiveData["currectPage"] = currectPage
		// arrange the pagenate
		storeLiveData["liveList"] = utils.Pagenator(liveList, currectPage, liveListSize)
	}

	resp, err := RespJSON{0, storeLiveData}.SetResponse()
	if err != nil {
		logger.ErrorLogger.Printf("Error on wrapping JSON resp, err: %v", err)
		errorMsg := utils.SetErrorMsg("Error on wrapping JSON resp")
		resp, _ := RespJSON{int(utils.InvalidAccessTokenCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusInternalServerError)
		return
	}
	ReturnResponse(w, resp, http.StatusOK)
}
