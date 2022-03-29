package db

import (
	"time"

	"cloud.google.com/go/firestore"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
)

type ProductObject struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Picture     string  `json:"picture"`
	StoreId     string  `json:"storeId"`
	CreateTime  string  `json:"createTime"`
	UpdateTime  string  `json:"updateTime"`
	IsDeleted   bool    `json:"isDeleted"`
}

func GetProductNewCount() int {
	dsnap, err := FireBaseClient.Collection(DbCollections["settings"]).Doc("productAutoIncrement").Get(DatabaseCtx)
	if err != nil {
		logger.WarningLogger.Printf("Cannot productAutoIncrement in settings. Error: %s", err)
		return -1
	}
	value := dsnap.Data()
	newProductId := value["number"].(int64) + 1
	logger.DebugLogger.Printf("Document data: %#v\n, %T, productId: %v", value["number"], value["number"], newProductId)
	// return strconv.Itoa(int(newProductId))
	return int(newProductId)
}
func UpdateProductCount(newProductCount int) error {
	_, err := FireBaseClient.Collection(DbCollections["settings"]).Doc("productAutoIncrement").Update(DatabaseCtx, []firestore.Update{
		{
			Path:  "number",
			Value: newProductCount,
		},
	})
	if err != nil {
		logger.WarningLogger.Printf("Error updating productAutoIncrement to %d, error: %s", newProductCount, err)
	}
	return err
}
func AddProductObj(productId string, userData map[string]interface{}) error {
	_, err := FireBaseClient.Collection(DbCollections["products"]).Doc(productId).Set(DatabaseCtx, userData)
	if err != nil {
		logger.WarningLogger.Printf("Error adding value. %s", err)
	}
	return err
}

func GetProductObj(productId string) map[string]interface{} {
	dsnap, err := FireBaseClient.Collection(DbCollections["products"]).Doc(productId).Get(DatabaseCtx)
	if err != nil {
		logger.WarningLogger.Printf("Cannot find product by productId. %s", productId)
		return nil
	}
	value := dsnap.Data()
	logger.DebugLogger.Printf("Document data: %#v\n", value)
	return value
}
func GetProductObj2(productId string) (ProductObject, error) {
	dsnap, err := FireBaseClient.Collection(DbCollections["products"]).Doc(productId).Get(DatabaseCtx)
	var p ProductObject
	if err != nil {
		logger.WarningLogger.Printf("Cannot find product by productId. %s", productId)
		return p, err
	}
	err = dsnap.DataTo(&p)
	if err != nil {
		logger.WarningLogger.Printf("Cannot find product by productId. %s", productId)
		return p, err
	}
	logger.DebugLogger.Printf("Document data: %#v\n", p)
	return p, nil
}
func UpdateTimeChange(productId string) error {
	nowTime := time.Now().UTC().Format(time.RFC3339)
	_, err := FireBaseClient.Collection(DbCollections["products"]).Doc(productId).Update(DatabaseCtx, []firestore.Update{
		{
			Path:  "updateTime",
			Value: nowTime,
		},
	})
	if err != nil {
		logger.WarningLogger.Printf("Error updating value on field updateTime. %s", err)
	}
	return err
}
func UpdateProductObj(productId string, fieldStr string, fieldValue interface{}) error {
	_, err := FireBaseClient.Collection(DbCollections["products"]).Doc(productId).Update(DatabaseCtx, []firestore.Update{
		{
			Path:  fieldStr,
			Value: fieldValue,
		},
	})
	if err != nil {
		logger.WarningLogger.Printf("Error updating value on field %s. %s", fieldStr, err)
	}
	return UpdateTimeChange(productId)
}
