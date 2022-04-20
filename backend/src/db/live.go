package db

import (
	"cloud.google.com/go/firestore"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
	"google.golang.org/api/iterator"
)

func GetLiveObj(livdeId string) map[string]interface{} {
	dsnap, err := FireBaseClient.Collection(DbCollections["lives"]).Doc(livdeId).Get(DatabaseCtx)
	if err != nil {
		logger.WarningLogger.Printf("Cannot find live by livdeId. %s", err)
		return nil
	}
	value := dsnap.Data()
	return value
}

func AddLiveObj(livdeId string, liveData map[string]interface{}) error {
	_, err := FireBaseClient.Collection(DbCollections["lives"]).Doc(livdeId).Set(DatabaseCtx, liveData)
	if err != nil {
		logger.ErrorLogger.Printf("Error adding value. %s", err)
	}
	return err
}

func GetLiveOrders(storeId string, page int) []map[string]interface{} {
	var orderList []map[string]interface{}
	iter := FireBaseClient.Collection(DbCollections["orders"]).Where("liveId", "==", storeId).OrderBy("createTime", firestore.Desc).Documents(DatabaseCtx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			logger.WarningLogger.Printf("GetStoreOrders on store %s Error iterating - Error %s", storeId, err)
		}
		orderList = append(orderList, doc.Data())
	}
	return orderList
}
