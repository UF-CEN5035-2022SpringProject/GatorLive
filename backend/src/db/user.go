package db

import (
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
)

func GetUserObj(userId string) {
	logger.DebugLogger.Printf("%v", FireBaseClient)

	dsnap, err := FireBaseClient.Collection(Collections["users"]).Doc(userId).Get(DatabaseCtx)
	if err != nil {
		// t.Errorf("Error retreiving value in TestDBGetUserObj")
		logger.ErrorLogger.Fatalf("Error retreiving value in TestDBGetUserObj. %s", err)
	}
	m := dsnap.Data()
	logger.DebugLogger.Printf("Document data: %#v\n", m)
}
