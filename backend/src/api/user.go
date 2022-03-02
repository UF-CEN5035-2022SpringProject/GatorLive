package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/db"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/utils"
	"github.com/gorilla/mux"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	b64 "encoding/base64"

	g "google.golang.org/api/oauth2/v2"
	youtube "google.golang.org/api/youtube/v3"
)

var (
	ClientID     string
	ClientSecret string
	RedirectURL  []string
)

type Code struct {
	Code string `json:"code"`
}
type WebStruct struct {
	Client_id     string
	Redirect_uris []string
	Client_secret string
}
type credential struct {
	Web WebStruct
}

type Response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}
type ResultSuccess struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	JwtToken string `json:"jwtToken"`
}
type ResultError struct {
	ErrorName string
}
type Profile struct {
	Name  string
	Email string
}

func GetUserProfile(accesstoken string) Profile {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	req.Header.Set("Authorization", "Bearer "+accesstoken)
	res, err := client.Do(req)

	if err != nil {
		logger.DebugLogger.Fatal(err)
		// log.Fatal(err)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		logger.DebugLogger.Fatalf("Unable to get Google profile: %v", err)
		// log.Fatalf("Unable to create YouTube service: %v", e)
	}
	// fmt.Println("profile:" + string(b))
	var profile Profile
	err = json.Unmarshal(b, &profile)
	if err != nil {
		logger.DebugLogger.Fatalf("Unable to decode Google profile: %v", err)
		// log.Fatalf("Unable to create YouTube service: %v", e)
	}
	return profile
}

func ReadCredential() {
	content, err := ioutil.ReadFile("./client_secret.json")
	if err != nil {
		logger.DebugLogger.Fatal(err)
	}
	var cre credential
	err = json.Unmarshal(content, &cre)
	if err != nil {
		logger.DebugLogger.Fatal(err)
	}
	ClientID = cre.Web.Client_id
	ClientSecret = cre.Web.Client_secret
	RedirectURL = cre.Web.Redirect_uris
}

func createJwtToken(userId string, userEmail string, nowTime string) string {
	// store newJwt in DB
	newJwtToken := "gst." + b64.StdEncoding.EncodeToString([]byte(utils.JwtPrefix+userEmail+userId)) + "_" + b64.StdEncoding.EncodeToString([]byte(nowTime))
	db.AddJwtToken(newJwtToken, userEmail, nowTime)
	return newJwtToken
}

// API ENTRYPOINT
func Login(w http.ResponseWriter, r *http.Request) {
	// setup config
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Scopes:       []string{youtube.YoutubeScope, g.UserinfoEmailScope, g.UserinfoProfileScope},

		Endpoint:    google.Endpoint,
		RedirectURL: RedirectURL[1],
		// Endpoint: oauth2.Endpoint{
		// 	AuthURL:  "https://provider.com/o/oauth2/auth",
		// 	TokenURL: "https://provider.com/o/oauth2/token",
		// },
	}

	// get code or assesstoken from http.request
	b, err := io.ReadAll(r.Body)
	logger.DebugLogger.Printf("request login body %s", b)

	if err != nil {
		logger.ErrorLogger.Printf("Unable to read login request body, err: %v", err)
		errorMsg := utils.SetErrorMsg("error occurs before google login")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusInternalServerError)
		return
	}

	var code Code
	err = json.Unmarshal(b, &code)
	logger.DebugLogger.Printf("request login code %s", code)
	if err != nil {
		logger.ErrorLogger.Printf("Unable to decode login request body, err: %v, google api code %s", err, code)
		errorMsg := utils.SetErrorMsg("error occurs before google login")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusInternalServerError)
		return
	}

	tok, err := conf.Exchange(ctx, code.Code)
	if err != nil {
		logger.ErrorLogger.Printf("Exchange token by code failed! err: %v", err)
		errorMsg := utils.SetErrorMsg("Exchange token by code failed!")
		resp, _ := RespJSON{int(utils.InvalidGoogleCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}

	client := conf.Client(ctx, tok)
	// service, e := youtube.New(client)
	_, err = youtube.New(client)
	if err != nil {
		logger.ErrorLogger.Printf("Unable to create YouTube service with token %v, err %v", tok, err)
		errorMsg := utils.SetErrorMsg("Exchange token by code failed!")
		resp, _ := RespJSON{int(utils.InvalidAccessTokenCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}
	profile := GetUserProfile(tok.AccessToken)

	tokenBytes, err := json.Marshal(tok)
	if err != nil {
		logger.ErrorLogger.Printf("Unable to decode token into byte, err: %v", err)
		errorMsg := utils.SetErrorMsg("error occurs after google login")
		resp, _ := RespJSON{int(utils.InvalidAccessTokenCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}

	tokenString := string(tokenBytes)
	// Flow: Check the user email
	//    - No email -> store and return the obj
	//    - Email -> update the token and return the obj
	userData := db.GetUserObj(profile.Email)
	if userData == nil {
		// create userId and assign JWT
		newUserCount := db.GetUserNewCount()
		newUserId := strconv.Itoa(newUserCount)
		logger.DebugLogger.Printf("New user, assign ID: %s", newUserId)

		// Add user Data
		nowTime := time.Now().UTC().Format(time.RFC3339)
		userObj := &db.UserObject{
			Id:          newUserId,
			Name:        profile.Name,
			Email:       profile.Email,
			JwtToken:    createJwtToken(newUserId, profile.Email, nowTime),
			AccessToken: tokenString,
			CreateTime:  nowTime,
			UpdateTime:  nowTime,
		}

		var convertMap map[string]interface{}
		userObjStr, _ := json.Marshal(userObj)
		json.Unmarshal(userObjStr, &convertMap)

		userData = convertMap

		db.AddUserObj(profile.Email, userData)
		db.UpdateUserCount(newUserCount)
	} else {
		db.UpdateUserObj(profile.Email, "accessToken", tokenString)
		userData = db.GetUserObj(profile.Email)
	}

	resp, err := RespJSON{0, userData}.SetResponse()
	if err != nil {
		logger.ErrorLogger.Printf("Error on wrapping JSON resp, err: %v", err)
		errorMsg := utils.SetErrorMsg("Error on wrapping JSON resp")
		resp, _ := RespJSON{int(utils.InvalidAccessTokenCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusInternalServerError)
		return
	}
	ReturnResponse(w, resp, http.StatusOK)
}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	// Depend on the action
	// 1. Get userInfo
	logger.DebugLogger.Println(r.Method)
	vars := mux.Vars(r)
	if r.Method == "GET" {
		fmt.Fprintf(w, "Get %v user info", vars["userId"])
		jwtToken := r.Header.Get("jwtToken")

		userEmail := db.MapJwtToken(jwtToken)["email"]
		userData := db.GetUserObj(userEmail.(string))

		if userData == nil {
			logger.ErrorLogger.Printf("Invalid JWT, unable to get user")
			errorMsg := utils.SetErrorMsg("Invalid JWT, unable to get user")
			resp, _ := RespJSON{int(utils.InvalidJwtTokenCode), errorMsg}.SetResponse()
			ReturnResponse(w, resp, http.StatusUnauthorized)
			return
		}

		if userData["id"] != vars["userId"] {
			logger.ErrorLogger.Printf("invald request")
			errorMsg := utils.SetErrorMsg("invald request")
			resp, _ := RespJSON{int(utils.InvalidJwtTokenCode), errorMsg}.SetResponse()
			ReturnResponse(w, resp, http.StatusBadRequest)
			return
		}

		resp, err := RespJSON{0, userData}.SetResponse()
		if err != nil {
			logger.ErrorLogger.Printf("Error on wrapping JSON resp, Error: %s", err)
		}

		ReturnResponse(w, resp, http.StatusOK)
		return
	} else if r.Method == "PUT" {
		fmt.Fprintf(w, "Update %v user info", vars["userId"])
	}
}
