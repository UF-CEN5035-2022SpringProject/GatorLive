package db

import (
	"cloud.google.com/go/firestore"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
)

type JwtObject struct {
	Email      string `json:"email"`
	JwtToken   string `json:"jwtToken"`
	CreateTime string `json: "createTime"`
}

type UserObject struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	JwtToken    string `json:"jwtToken"`
	AccessToken string `json:"accessToken"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}

/*** JWT functions ***/
func MapJwtToken(jwtToken string) map[string]interface{} {
	logger.DebugLogger.Printf("Find jwtObj by token %s", jwtToken)
	dsnap, err := FireBaseClient.Collection(DbCollections["jwtTokenMap"]).Doc(jwtToken).Get(DatabaseCtx)
	if err != nil {
		logger.WarningLogger.Printf("Cannot find user jwtToken (%s). %s", jwtToken, err)
		return nil
	}
	value := dsnap.Data()
	logger.DebugLogger.Printf("Document data: %#v\n", value)
	return value
}

func AddJwtToken(jwtToken string, userEmail string, nowTime string) error {
	newJwtObject := JwtObject{
		Email:      userEmail,
		JwtToken:   jwtToken,
		CreateTime: nowTime,
	}
	_, err := FireBaseClient.Collection(DbCollections["jwtTokenMap"]).Doc(jwtToken).Set(DatabaseCtx, newJwtObject)
	if err != nil {
		logger.WarningLogger.Printf("Error adding value. %s", err)
	}
	return err
}

/*** User functions ***/
func GetUserNewCount() int {
	dsnap, err := FireBaseClient.Collection(DbCollections["settings"]).Doc("userAutoIncrement").Get(DatabaseCtx)
	if err != nil {
		logger.WarningLogger.Printf("Cannot userAutoIncrement in settings. Error: %s", err)
		return -1
	}
	value := dsnap.Data()
	newUserId := value["number"].(int64) + 1
	logger.DebugLogger.Printf("Document data: %#v\n, %T, newUserId: %v", value["number"], value["number"], newUserId)
	// return strconv.Itoa(int(newUserId))
	return int(newUserId)
}

func GetUserObj(userEmail string) map[string]interface{} {
	dsnap, err := FireBaseClient.Collection(DbCollections["users"]).Doc(userEmail).Get(DatabaseCtx)
	if err != nil {
		logger.WarningLogger.Printf("Cannot find user by email. %s", err)
		return nil
	}
	value := dsnap.Data()
	logger.DebugLogger.Printf("Document data: %#v\n", value)
	return value
}

func AddUserCount(newUserCount int) error {
	_, err := FireBaseClient.Collection(DbCollections["settings"]).Doc("userAutoIncrement").Set(DatabaseCtx, newUserCount)
	if err != nil {
		logger.WarningLogger.Printf("Error adding value. %s", err)
	}
	return err
}

func AddUserObj(userEmail string, userData map[string]interface{}) error {
	_, err := FireBaseClient.Collection(DbCollections["users"]).Doc(userEmail).Set(DatabaseCtx, userData)
	if err != nil {
		logger.WarningLogger.Printf("Error adding value. %s", err)
	}
	return err
}

func UpdateUserObj(userEmail string, fieldStr string, fieldValue interface{}) error {
	_, err := FireBaseClient.Collection(DbCollections["users"]).Doc(userEmail).Update(DatabaseCtx, []firestore.Update{
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
