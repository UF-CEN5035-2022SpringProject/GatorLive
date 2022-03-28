package api

import (
	"encoding/json"
	"fmt"
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

type ProductCreateObject struct {
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
	Picture     string `json:"picture"`
	StoreId     string `json:"StoreId"`
}

func ProductCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logger.ErrorLogger.Printf("ProductCreate with wrong method: %v", r.Method)
		errorMsg := utils.SetErrorMsg("wrong method")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}

	b, err := io.ReadAll(r.Body)

	if err != nil {
		logger.ErrorLogger.Printf("Unable to read ProductCreate request body, err: %v", err)
		errorMsg := utils.SetErrorMsg("error on ProductCreate")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusInternalServerError)
		return
	}

	var product ProductCreateObject
	err = json.Unmarshal(b, &product)

	if err != nil {
		logger.ErrorLogger.Printf("Unable to decode ProductCreate request body, err: %v", err)
		errorMsg := utils.SetErrorMsg("error occurs before ProductCreate")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusInternalServerError)
		return
	}

	userData := gorillaContext.Get(r, "userData").(map[string]interface{})
	storeObj := db.GetStoreObj(product.StoreId)
	storeUserId := fmt.Sprintf("%v", storeObj["userId"])
	JwtUserId := fmt.Sprintf("%v", userData["id"])

	if storeUserId != JwtUserId {
		logger.ErrorLogger.Printf("jwtToken and storeId not match %v != %v", JwtUserId, storeUserId)
		errorMsg := utils.SetErrorMsg("jwtToken and storeId not match")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusForbidden)
		return
	}

	newProductCount := db.GetProductNewCount()
	newProductId := "Product-" + strconv.Itoa(newProductCount)
	nowTime := time.Now().UTC().Format(time.RFC3339)
	productObj := &db.ProductObject{
		Id:          newProductId,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Quantity:    product.Quantity,
		Picture:     product.Picture,
		StoreId:     product.StoreId,
		IsDeleted:   false,
		CreateTime:  nowTime,
		UpdateTime:  nowTime,
	}

	var convertMap map[string]interface{}
	userObjStr, _ := json.Marshal(productObj)
	json.Unmarshal(userObjStr, &convertMap)

	db.AddProductObj(newProductId, convertMap)
	db.UpdateProductCount(newProductCount)

	resp, err := RespJSON{0, convertMap}.SetResponse()
	if err != nil {
		logger.ErrorLogger.Printf("Error on wrapping JSON resp, err: %v", err)
		errorMsg := utils.SetErrorMsg("Error on wrapping JSON resp")
		resp, _ := RespJSON{int(utils.MissingParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusInternalServerError)
		return
	}
	ReturnResponse(w, resp, http.StatusOK)
}
func ProductGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	productObj := db.GetProductObj(productId)

	if productObj == nil {
		logger.ErrorLogger.Printf("Unable to get product, id: %v", productId)
		errorMsg := utils.SetErrorMsg("Unable to get product")
		resp, _ := RespJSON{int(utils.UnableToGetDbObj), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusNotFound)
		return
	}

	resp, err := RespJSON{0, productObj}.SetResponse()
	if err != nil {
		logger.ErrorLogger.Printf("Error on wrapping JSON resp, Error: %s", err)
	}
	ReturnResponse(w, resp, http.StatusOK)
}

func ProductPut(w http.ResponseWriter, r *http.Request) {

}
func ProductPost(w http.ResponseWriter, r *http.Request) {

}

func ProductDelete(w http.ResponseWriter, r *http.Request) {

}
func ProductRESTFUL(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ProductGet(w, r)
	} else if r.Method == http.MethodPut {
		ProductPut(w, r)
	} else if r.Method == http.MethodPost {
		ProductPost(w, r)
	} else if r.Method == http.MethodDelete {
		ProductDelete(w, r)
	}
}
