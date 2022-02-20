package api

import (
	"context"
	"fmt"

	// "encoding/json"
	// "fmt"
	"encoding/json"
	"io"
	"time"

	// "io/ioutil"
	"net/http"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/db"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"

	// "github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"

	// "golang.org/x/oauth2/google"
	// g "google.golang.org/api/oauth2/v2"
	youtube "google.golang.org/api/youtube/v3"
)

type Title struct {
	Title string `json:"title"`
}

// func token(accessToken string) (*oauth2.Token, error) {
// 	return &oauth2.Token{
// 		AccessToken: accessToken,
// 		TokenType:   "Bearer",
// 	}, nil
// }
func getStream(service *youtube.Service) *youtube.LiveStream {
	list := service.LiveStreams.List([]string{"id", "cdn"})
	list = list.Mine(true)
	rList, err := list.Do()
	if err != nil {
		logger.DebugLogger.Panicf("Error making YouTube API call list: %v\n", err)
	}
	if len(rList.Items) != 0 {
		return rList.Items[0]
	}
	newStream := &youtube.LiveStream{
		Snippet: &youtube.LiveStreamSnippet{
			Title: "GatorStore stream",
		},
		Cdn: &youtube.CdnSettings{
			FrameRate:     "60fps",
			IngestionType: "rtmp",
			Resolution:    "1080p",
		},
		ContentDetails: &youtube.LiveStreamContentDetails{
			IsReusable: true,
		},
	}
	stream := service.LiveStreams.Insert([]string{"snippet", "cdn", "contentDetails", "status"}, newStream)
	newStream, err = stream.Do()
	if err != nil {
		logger.DebugLogger.Panicf("Error making YouTube API call stream: %v\n", err)
	}
	return newStream
}
func bind(service *youtube.Service, live *youtube.LiveBroadcast, stream *youtube.LiveStream) {
	bindS := service.LiveBroadcasts.Bind(live.Id, []string{"snippet", "status"})
	bindS = bindS.StreamId(stream.Id)
	_, err := bindS.Do()
	if err != nil {
		logger.DebugLogger.Panicf("Error making YouTube API call bind: %v\n", err)
	}
}
func verify(jwtToken string, storeid string) string {
	return "1"
}
func GetEmail(jwtToken string) map[string]interface{} {
	dsnap, err := db.FireBaseClient.Collection("jwtTokenMap").Doc(jwtToken).Get(db.DatabaseCtx)
	if err != nil {
		logger.WarningLogger.Printf("Cannot find user by email. %s", err)
		return nil
	}
	value := dsnap.Data()
	logger.DebugLogger.Printf("Document data: %#v\n", value)

	return value
}
func CreateLivebroadcast(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storeId := vars["storeId"]
	jwtToken := r.Header.Get("Authorization")

	// TODO verify
	if verify(jwtToken, storeId) == "" {
		return
	}

	b, err := io.ReadAll(r.Body)
	// logger.DebugLogger.Printf("request livestream create body %s", b)

	if err != nil {
		logger.DebugLogger.Panicf("Unable to read livestream create req: %v", err)
	}
	var title Title

	err = json.Unmarshal(b, &title)

	if err != nil {
		logger.DebugLogger.Panicf("Unable to decode livestream create req: %v, code %s", err, jwtToken)
		// log.Fatalf("Unable to create YouTube service: %v", e)
	}

	emailObj := GetEmail(jwtToken)
	email := fmt.Sprintf("%v", emailObj["Email"])
	userProfile := db.GetUserObj(email)
	tokenByte := []byte(fmt.Sprintf("%v", userProfile["accessToken"]))

	var accessToken oauth2.Token
	err = json.Unmarshal(tokenByte, &accessToken)
	if err != nil {
		logger.DebugLogger.Panicf("Unable to decode accessToken: %v", err)
		// log.Fatalf("Unable to create YouTube service: %v", e)
	}
	ctx := context.Background()
	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&accessToken))
	service, e := youtube.New(client)
	if e != nil {
		logger.DebugLogger.Panicf("Unable to create YouTube service: %v", e)
	}
	createTime := time.Now()
	startTime := createTime.Add(time.Minute * 10)
	endTime := startTime.Add((time.Hour * 24))
	newLive := &youtube.LiveBroadcast{
		Snippet: &youtube.LiveBroadcastSnippet{
			Title:              storeId + "-" + title.Title,
			ScheduledStartTime: startTime.UTC().Format(time.RFC3339),
			ScheduledEndTime:   endTime.UTC().Format(time.RFC3339),
		},
		Status: &youtube.LiveBroadcastStatus{
			PrivacyStatus: "unlisted",
		},
	}
	// newLive.Snippet. = []string{"test","api"}
	call := service.LiveBroadcasts.Insert([]string{"snippet", "status"}, newLive)
	newLive, err = call.Do()
	if err != nil {
		logger.DebugLogger.Panicf("Error making YouTube API call: %v\n", err)
	}

	stream := getStream(service)
	bind(service, newLive, stream)

	response := make(map[string]interface{})
	response["id"] = newLive.Id
	response["title"] = newLive.Snippet.Title
	response["streamKey"] = stream.Cdn.IngestionInfo.StreamName
	response["streamUrl"] = stream.Cdn.IngestionInfo.IngestionAddress
	response["createTime"] = createTime.UTC().Format(time.RFC3339)
	response["updateTime"] = createTime.UTC().Format(time.RFC3339)

	resp, err := JsonResponse(response, 0)

	if err != nil {
		logger.ErrorLogger.Fatalf("Error on wrapping JSON resp %s", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
