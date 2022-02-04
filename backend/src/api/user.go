package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/db"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
	"github.com/gorilla/mux"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

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
func Login(w http.ResponseWriter, r *http.Request) {
	// TODO @chouhy

	// setup config
	logger.DebugLogger.Println("User ___ Login")
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Scopes:       []string{youtube.YoutubeScope, g.UserinfoEmailScope, g.UserinfoProfileScope},

		Endpoint:    google.Endpoint,
		RedirectURL: RedirectURL[0],
		// Endpoint: oauth2.Endpoint{
		// 	AuthURL:  "https://provider.com/o/oauth2/auth",
		// 	TokenURL: "https://provider.com/o/oauth2/token",
		// },
	}

	// get code or assesstoken from http.request
	b, err := io.ReadAll(r.Body)
	if err != nil {
		logger.DebugLogger.Fatalf("Unable to read login req: %v", err)
		// log.Fatalf("Unable to create YouTube service: %v", e)
	}
	var code Code
	err = json.Unmarshal(b, &code)
	if err != nil {
		logger.DebugLogger.Fatalf("Unable to decode login req: %v", err)
		// log.Fatalf("Unable to create YouTube service: %v", e)
	}
	tok, err := conf.Exchange(ctx, code.Code)

	if err != nil {
		logger.DebugLogger.Panic(err)
		// log.Fatal(err)
	}

	client := conf.Client(ctx, tok)
	// service, e := youtube.New(client)
	_, err = youtube.New(client)
	if err != nil {
		logger.DebugLogger.Fatalf("Unable to create YouTube service: %v", err)
	}
	profile := GetUserProfile(tok.AccessToken)

	// Flow: Check the user email
	//    - No email -> store and return the obj
	//    - Email -> update the token and return the obj
	userData := db.GetUserObj(profile.Email)
	if userData == nil {
		// Add user Data
		userData = make(map[string]interface{})
		userData["id"] = "113024"
		userData["name"] = profile.Name
		userData["email"] = profile.Email
		userData["jwtToken"] = "gatorStore_qeqweiop122133"
		userData["accessToken"] = tok.AccessToken
		db.AddUserObj(profile.Email, userData)
	} else {
		db.UpdateUserObj(profile.Email, "accessToken", tok.AccessToken)
		userData = db.GetUserObj(profile.Email)
	}

	resp, err := JsonResponse(userData, 0)
	if err != nil {
		logger.ErrorLogger.Fatalf("Error on wrapping JSON resp %s", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	// Depend on the action
	// 1. Get userInfo
	logger.DebugLogger.Println(r.Method)
	vars := mux.Vars(r)
	if r.Method == "GET" {
		fmt.Fprintf(w, "Get %v user info", vars["userId"])
		value := db.GetUserObj(vars["userId"])
		resp, err := JsonResponse(value, 0)
		if err != nil {
			logger.ErrorLogger.Fatalf("Error on wrapping JSON resp %s", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)

	} else if r.Method == "PUT" {
		fmt.Fprintf(w, "Update %v user info", vars["userId"])
	}
}
