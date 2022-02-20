package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/db"
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
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Add authentication
		token := r.Header.Get("Authorization")
		logger.DebugLogger.Printf("Authorization header token %s\n", token)

		if token == "" {
			logger.WarningLogger.Printf("Authorization empty token %s\n", token)
			http.Error(w, "UnAuthorized", http.StatusUnauthorized)
		}

		jwtMap := db.MapJwtToken(token)

		if user, found := jwtMap[token]; found {
			// We found the token in our map
			logger.DebugLogger.Printf("Authenticated user %s\n", user)
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			http.Error(w, "UnAuthorized", http.StatusUnauthorized)
		}
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		logger.DebugLogger.Println(r.RequestURI)
		logger.DebugLogger.Println(r.URL.RawQuery)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
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

	jsonResponse, err := json.Marshal(respObj)
	if err != nil {
		fmt.Println("Unable to encode JSON")
	}
	return jsonResponse, err
}
