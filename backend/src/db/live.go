package db

import "github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"

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
