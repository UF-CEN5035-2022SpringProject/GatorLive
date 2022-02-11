package db

import (
	"cloud.google.com/go/firestore"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
)

func GetUserObj(userEmail string) map[string]interface{} {
	dsnap, err := FireBaseClient.Collection(Collections["users"]).Doc(userEmail).Get(DatabaseCtx)
	if err != nil {
		logger.WarningLogger.Printf("Cannot find user by email. %s", err)
		return nil
	}
	value := dsnap.Data()
	logger.DebugLogger.Printf("Document data: %#v\n", value)

	return value
}

func AddUserObj(userEmail string, userData map[string]interface{}) error {
	_, err := FireBaseClient.Collection("users").Doc(userEmail).Set(DatabaseCtx, userData)
	if err != nil {
		logger.WarningLogger.Printf("Error adding value. %s", err)
	}
	return err
}

func UpdateUserObj(userEmail string, fieldStr string, fieldValue interface{}) error {
	_, err := FireBaseClient.Collection("users").Doc(userEmail).Update(DatabaseCtx, []firestore.Update{
		{
			Path:  fieldStr,
			Value: fieldValue,
		},
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		logger.WarningLogger.Printf("Error updating value on field %s. %s", fieldStr, err)
	}
	return err
}
