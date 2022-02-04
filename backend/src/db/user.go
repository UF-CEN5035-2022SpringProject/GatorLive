package db

import (
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
)

func GetUserObj(userId string) map[string]interface{} {
	dsnap, err := FireBaseClient.Collection(Collections["users"]).Doc(userId).Get(DatabaseCtx)
	if err != nil {
		logger.ErrorLogger.Fatalf("Error retreiving value in GetUserObj. %s", err)
	}
	value := dsnap.Data()
	logger.DebugLogger.Printf("Document data: %#v\n", value)

	return value
}
