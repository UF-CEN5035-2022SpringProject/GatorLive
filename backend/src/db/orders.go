package db

import (
	"cloud.google.com/go/firestore"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
)

type OrderObject struct {
	Id         string  `json:"id"`
	UserId     string  `json:"userId"`
	Subtotal   float64 `json:"subtotal"`
	LiveId     string  `json:"liveId"`
	Quantity   int     `json:"quantity"`
	ProductId  string  `json:"productId"`
	CreateTime string  `json:"createTime"`
}

func GetOrderNewCount() int {
	dsnap, err := FireBaseClient.Collection(DbCollections["settings"]).Doc("orderAutoIncrement").Get(DatabaseCtx)
	if err != nil {
		logger.WarningLogger.Printf("Cannot orderAutoIncrement in settings. Error: %s", err)
		return -1
	}
	value := dsnap.Data()
	newProductId := value["number"].(int64) + 1
	logger.DebugLogger.Printf("Document data: %#v\n, %T, orderId: %v", value["number"], value["number"], newProductId)
	// return strconv.Itoa(int(newProductId))
	return int(newProductId)
}
func UpdateOrderCount(newOrderCount int) error {
	_, err := FireBaseClient.Collection(DbCollections["settings"]).Doc("orderAutoIncrement").Update(DatabaseCtx, []firestore.Update{
		{
			Path:  "number",
			Value: newOrderCount,
		},
	})
	if err != nil {
		logger.WarningLogger.Printf("Error updating orderAutoIncrement to %d, error: %s", newOrderCount, err)
	}
	return err
}
func AddOrderObj(orderId string, userData map[string]interface{}) error {
	_, err := FireBaseClient.Collection(DbCollections["orders"]).Doc(orderId).Set(DatabaseCtx, userData)
	if err != nil {
		logger.WarningLogger.Printf("Error adding value. %s", err)
	}
	return err
}
func GetOrderObj(orderId string) map[string]interface{} {
	dsnap, err := FireBaseClient.Collection(DbCollections["orders"]).Doc(orderId).Get(DatabaseCtx)
	if err != nil {
		logger.WarningLogger.Printf("Cannot find order by orderId. %s", orderId)
		return nil
	}
	value := dsnap.Data()
	logger.DebugLogger.Printf("Document data: %#v\n", value)
	return value
}
