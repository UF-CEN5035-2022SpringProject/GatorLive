package db

import (
	firestore "cloud.google.com/go/firestore"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
	"google.golang.org/api/iterator"
)

/*** User functions ***/
func GetStoreNewCount() int {
	dsnap, err := FireBaseClient.Collection(DbCollections["settings"]).Doc("storeAutoIncrement").Get(DatabaseCtx)
	if err != nil {
		logger.WarningLogger.Printf("Cannot storeAutoIncrement in settings. Error: %s", err)
		return -1
	}
	value := dsnap.Data()
	newStoreId := value["number"].(int64) + 1
	logger.DebugLogger.Printf("Document data: %#v\n, %T, newStoreId: %v", value["number"], value["number"], newStoreId)
	// return strconv.Itoa(int(newUserId))
	return int(newStoreId)
}

func UpdateStoreCount(newStoreCount int) error {
	_, err := FireBaseClient.Collection(DbCollections["settings"]).Doc("storeAutoIncrement").Update(DatabaseCtx, []firestore.Update{
		{
			Path:  "number",
			Value: newStoreCount,
		},
	})
	if err != nil {
		logger.WarningLogger.Printf("Error updating storeAutoIncrement to %d, error: %s", newStoreCount, err)
	}
	return err
}

func AddStoreObj(storeId string, storeData map[string]interface{}) error {
	_, err := FireBaseClient.Collection(DbCollections["stores"]).Doc(storeId).Set(DatabaseCtx, storeData)
	if err != nil {
		logger.WarningLogger.Printf("Error adding value. %s", err)
	}
	return err
}

func GetStoreObj(storeId string) map[string]interface{} {
	dsnap, err := FireBaseClient.Collection(DbCollections["stores"]).Doc(storeId).Get(DatabaseCtx)
	if err != nil {
		logger.WarningLogger.Printf("Cannot find store by storeId. %s", storeId)
		return nil
	}
	value := dsnap.Data()
	logger.DebugLogger.Printf("Document data: %#v\n", value)
	return value
}

func GetStoreProducts(storeId string, page int) []map[string]interface{} {
	logger.DebugLogger.Printf("GetStoreProducts from storeId: %s", storeId)
	var productList []map[string]interface{}
	// OrderBy("id", firestore.Asc)
	iter := FireBaseClient.Collection(DbCollections["products"]).Where("StoreId", "==", storeId).Documents(DatabaseCtx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			logger.WarningLogger.Printf("Error iterating. %s", storeId)
		}
		productList = append(productList, doc.Data())
	}

	return productList
}
