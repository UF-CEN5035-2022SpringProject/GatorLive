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
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/utils"

	// "github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
	gorillaContext "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"

	// "golang.org/x/oauth2/google"
	// g "google.golang.org/api/oauth2/v2"
	youtube "google.golang.org/api/youtube/v3"
)

type Title struct {
	Title         string   `json:"title"`
	ProductIdList []string `json:"productIdList"`
}

type LiveId struct {
	LiveId string `json:"liveId"`
}

// func token(accessToken string) (*oauth2.Token, error) {
// 	return &oauth2.Token{
// 		AccessToken: accessToken,
// 		TokenType:   "Bearer",
// 	}, nil
// }
func getStream(service *youtube.Service) (*youtube.LiveStream, error) {
	list := service.LiveStreams.List([]string{"id", "cdn"})
	list = list.Mine(true)
	rList, err := list.Do()
	if err != nil {
		// logger.ErrorLogger.Printf("Error making YouTube API call list: %v\n", err)
		return nil, err
	}
	if len(rList.Items) != 0 {
		return rList.Items[0], nil
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
		// logger.ErrorLogger.Printf("Error making YouTube API call stream: %v\n", err)
		return nil, err
	}
	return newStream, nil
}
func bind(service *youtube.Service, live *youtube.LiveBroadcast, stream *youtube.LiveStream) error {
	bindS := service.LiveBroadcasts.Bind(live.Id, []string{"snippet", "status"})
	bindS = bindS.StreamId(stream.Id)
	_, err := bindS.Do()
	if err != nil {
		return err
	}
	return nil
}

func GetEmail(jwtToken string) map[string]interface{} {
	dsnap, err := db.FireBaseClient.Collection("jwtTokenMap").Doc(jwtToken).Get(db.DatabaseCtx)
	if err != nil {
		logger.WarningLogger.Printf("Cannot find user by email. %s", err)
		return nil
	}
	value := dsnap.Data()
	logger.ErrorLogger.Printf("Document data: %#v\n", value)

	return value
}
func CreateLivebroadcast(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storeId := vars["storeId"]
	jwtToken := r.Header.Get("Authorization")

	// verify
	userData := gorillaContext.Get(r, "userData").(map[string]interface{})

	storeObj := db.GetStoreObj(storeId)
	if storeObj == nil {
		logger.ErrorLogger.Printf("invald request, unable to get store")
		errorMsg := utils.SetErrorMsg("invald request, unable to get store")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}

	userId := userData["id"].(string)
	if storeObj["userId"] != userId {
		logger.ErrorLogger.Printf("invald request, permission denied")
		errorMsg := utils.SetErrorMsg("invald request, permission denied")
		resp, _ := RespJSON{int(utils.InvalidJwtTokenCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusForbidden)
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil || jwtToken == "" {
		logger.ErrorLogger.Printf("Unable to read livestream create req: %v", err)
		errorMsg := utils.SetErrorMsg("Unable to read livestream create req")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}

	var title Title
	err = json.Unmarshal(b, &title)

	if err != nil {
		logger.ErrorLogger.Printf("Unable to decode livestream create req: %v, code %s", err, jwtToken)
		errorMsg := utils.SetErrorMsg("Unable to decode livestream create req")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}

	emailObj := db.MapJwtToken(jwtToken)
	if emailObj == nil {
		logger.ErrorLogger.Printf("Invalid JwtToken")
		errorMsg := utils.SetErrorMsg("Invalid JwtToken")
		resp, _ := RespJSON{int(utils.InvalidJwtTokenCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusUnauthorized)
		return
	}
	email := fmt.Sprintf("%v", emailObj["Email"])
	userProfile := db.GetUserObj(email)
	tokenByte := []byte(fmt.Sprintf("%v", userProfile["accessToken"]))

	var accessToken oauth2.Token
	err = json.Unmarshal(tokenByte, &accessToken)
	if err != nil {
		logger.ErrorLogger.Printf("Unable to decode accessToken: %v", err)
		errorMsg := utils.SetErrorMsg("Unable to decode accessToken")
		resp, _ := RespJSON{int(utils.InvalidAccessTokenCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusUnauthorized)
		return
	}

	ctx := context.Background()
	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&accessToken))
	service, e := youtube.New(client)
	if e != nil {
		logger.ErrorLogger.Printf("Unable to create YouTube service: %v", e)
		errorMsg := utils.SetErrorMsg("Unable to create YouTube service")
		resp, _ := RespJSON{int(utils.InvalidAccessTokenCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusUnauthorized)
		return
	}
	createTime := time.Now()
	// startTime := createTime.Add(time.Minute * 10)
	// endTime := startTime.Add((time.Hour * 24))

	startTime := createTime
	endTime := startTime.Add((time.Hour * 24))

	newLive := &youtube.LiveBroadcast{
		ContentDetails: &youtube.LiveBroadcastContentDetails{
			EnableAutoStart:    true,
			EnableAutoStop:     true,
			LatencyPreference:  "ultraLow",
			ClosedCaptionsType: "closedCaptionsDisabled",
		},
		Snippet: &youtube.LiveBroadcastSnippet{
			Title:              storeId + "-" + title.Title,
			ScheduledStartTime: startTime.UTC().Format(time.RFC3339),
			ScheduledEndTime:   endTime.UTC().Format(time.RFC3339),
		},
		Status: &youtube.LiveBroadcastStatus{
			PrivacyStatus:           "unlisted",
			SelfDeclaredMadeForKids: false,
		},
	}
	// newLive.Snippet. = []string{"test","api"}
	call := service.LiveBroadcasts.Insert([]string{"snippet", "contentDetails", "status"}, newLive)
	newLive, err = call.Do()
	if err != nil {
		logger.ErrorLogger.Printf("Error making YouTube API call: %v\n", err)
		errorMsg := utils.SetErrorMsg("Error making YouTube API call")
		resp, _ := RespJSON{int(utils.InvalidAccessTokenCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusUnauthorized)
		return
	}

	stream, err := getStream(service)
	if err != nil {
		logger.ErrorLogger.Printf("Error make YouTube API get/create Stream: %v\n", err)
		errorMsg := utils.SetErrorMsg("Error make YouTube API get/create Stream")
		resp, _ := RespJSON{int(utils.InvalidAccessTokenCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusUnauthorized)
		return
	}
	err = bind(service, newLive, stream)
	if err != nil {
		logger.ErrorLogger.Printf("Error binding YouTube broadcast and stream: %v\n", err)
		errorMsg := utils.SetErrorMsg("Error binding YouTube broadcast and stream")
		resp, _ := RespJSON{int(utils.InvalidAccessTokenCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusUnauthorized)
		return
	}

	liveObj := make(map[string]interface{})
	liveObj["id"] = newLive.Id
	liveObj["title"] = newLive.Snippet.Title
	liveObj["streamKey"] = stream.Cdn.IngestionInfo.StreamName
	liveObj["streamUrl"] = stream.Cdn.IngestionInfo.IngestionAddress
	liveObj["createTime"] = createTime.UTC().Format(time.RFC3339)
	liveObj["updateTime"] = createTime.UTC().Format(time.RFC3339)
	liveObj["embedHTML"] = newLive.ContentDetails.MonitorStream.EmbedHtml
	liveObj["productList"] = title.ProductIdList
	liveObj["storeId"] = storeId

	if db.GetLiveObj(newLive.Id) == nil {
		db.AddLiveObj(newLive.Id, liveObj)
	} else {
		// TODO: Update Live Obj
	}

	productObjList := make([]map[string]interface{}, len(title.ProductIdList))
	for index := 0; index < len(title.ProductIdList); index++ {
		productObjList[index] = db.GetProductObj(title.ProductIdList[index])
	}

	liveObj["productList"] = productObjList

	resp, _ := RespJSON{0, liveObj}.SetResponse()
	ReturnResponse(w, resp, http.StatusOK)
}

func GetLiveStream(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		logger.ErrorLogger.Printf("Unable to read getlivestream req: %v", err)
		errorMsg := utils.SetErrorMsg("Unable to read getlivestream req")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}

	var liveId LiveId
	err = json.Unmarshal(b, &liveId)

	if err != nil {
		logger.ErrorLogger.Printf("Unable to decode getlivestream req: %v", err)
		errorMsg := utils.SetErrorMsg("Unable to decode getlivestream req")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}

	liveObj := db.GetLiveObj(liveId.LiveId)

	if liveObj == nil {
		logger.ErrorLogger.Printf("invald request, unable to get livestream")
		errorMsg := utils.SetErrorMsg("invald request, unable to get livestream")
		resp, _ := RespJSON{int(utils.InvalidParamsCode), errorMsg}.SetResponse()
		ReturnResponse(w, resp, http.StatusBadRequest)
		return
	}

	detail := r.URL.Query().Get("detail")
	if detail == "" {
		detail = "true"
	}
	if detail == "true" {
		productIdList := liveObj["productList"].([]interface{})
		productObjList := make([]map[string]interface{}, len(productIdList))
		for index := 0; index < len(productIdList); index++ {
			productObjList[index] = db.GetProductObj(productIdList[index].(string))
		}

		liveObj["productList"] = productObjList

	} else {
		liveObj["productList"] = nil
	}
	liveObj["streamKey"] = ""
	liveObj["streamUrl"] = ""
	resp, _ := RespJSON{0, liveObj}.SetResponse()
	ReturnResponse(w, resp, http.StatusOK)

}
