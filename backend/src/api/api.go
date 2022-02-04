package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
)

func HeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		w.Header().Set("Content-Type", "application/json")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func CrossAllowMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO: Add authentication
		next.ServeHTTP(w, r)
	})
}

type respJSON struct {
	Status int                    `json:"status"`
	Result map[string]interface{} `json:"result"`
}

func JsonResponse(result map[string]interface{}, errorCode int) ([]byte, error) {
	respObj := &respJSON{
		Status: errorCode,
		Result: result,
	}

	logger.DebugLogger.Printf("Json responseObj: %#v\n", respObj)
	jsonResponse, err := json.Marshal(respObj)
	if err != nil {
		fmt.Println("Unable to encode JSON")
	}
	logger.DebugLogger.Printf("Json Resp: %#v\n", jsonResponse)
	return jsonResponse, err
}
