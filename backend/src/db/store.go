package db

import "github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"

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
