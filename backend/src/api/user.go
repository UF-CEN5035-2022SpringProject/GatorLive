package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// "log"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
	"github.com/gorilla/mux"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	youtube "google.golang.org/api/youtube/v3"
)

var (
	ClientID     string
	ClientSecret string
	RedirectURL  []string
)

type WebStruct struct {
	Client_id     string
	Redirect_uris []string
	Client_secret string
}
type credential struct {
	Web WebStruct
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
		Scopes:       []string{youtube.YoutubeScope},

		Endpoint:    google.Endpoint,
		RedirectURL: RedirectURL[0],
		// Endpoint: oauth2.Endpoint{
		// 	AuthURL:  "https://provider.com/o/oauth2/auth",
		// 	TokenURL: "https://provider.com/o/oauth2/token",
		// },
	}
	// get code or assesstoken from http.request
	keys, ok := r.URL.Query()["code"]

	if !ok || len(keys[0]) < 1 {
		logger.ErrorLogger.Panic("Url Param 'key' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	code := keys[0]
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		logger.DebugLogger.Fatal(err)
		// log.Fatal(err)
	}
	fmt.Println("TOKEN: " + tok.AccessToken + " " + tok.TokenType)
	// https://youtube.googleapis.com/youtube/v3/liveBroadcasts?part=snippet%2CcontentDetails%2Cstatus&key=AIzaSyA9rodcA1a-K6QNBMWgBXmNw2zkUsP7WNg
	client := conf.Client(ctx, tok)
	// service, e := youtube.New(client)
	_, e := youtube.New(client)
	if e != nil {
		logger.DebugLogger.Fatalf("Unable to create YouTube service: %v", e)
		// log.Fatalf("Unable to create YouTube service: %v", e)
	}
}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	// Depend on the action
	// 1. Get userInfo
	logger.DebugLogger.Println(r.Method)
	vars := mux.Vars(r)
	if r.Method == "GET" {
		fmt.Fprintf(w, "Get %v user info", vars["userId"])
	} else if r.Method == "PUT" {
		fmt.Fprintf(w, "Update %v user info", vars["userId"])
	}
}
