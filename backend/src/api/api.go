package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/UF-CEN5035-2022SpringProject/GatorStore/db"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/utils"

	gorillaContext "github.com/gorilla/context"
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
		w.Header().Set("Access-Control-Request-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		logger.DebugLogger.Printf("Authorization header token %s\n", token)

		if token == "" {
			logger.WarningLogger.Printf("Authorization empty token %s\n", token)
			errorMsg := utils.SetErrorMsg("Empty JWT")
			resp, _ := RespJSON{int(utils.MissingJwtTokenCode), errorMsg}.SetResponse()
			ReturnResponse(w, resp, http.StatusUnauthorized)
			return
		}

		jwtMap := db.MapJwtToken(token)
		logger.DebugLogger.Printf("jwtMap %v\n", jwtMap)
		if jwtMap != nil {
			// Pass down the request to the next middleware (or final handler)
			userData := db.GetUserObj(jwtMap["Email"].(string))
			if userData == nil {
				logger.ErrorLogger.Printf("Invalid JWT, user obj not found")
				errorMsg := utils.SetErrorMsg("Invalid JWT, user obj not found")
				resp, _ := RespJSON{int(utils.InvalidJwtTokenCode), errorMsg}.SetResponse()
				ReturnResponse(w, resp, http.StatusUnauthorized)
				return
			}
			gorillaContext.Set(r, "userData", userData)
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			logger.ErrorLogger.Printf("Invalid JWT, not found")
			errorMsg := utils.SetErrorMsg("Invalid JWT, not found")
			resp, _ := RespJSON{int(utils.MissingJwtTokenCode), errorMsg}.SetResponse()
			ReturnResponse(w, resp, http.StatusUnauthorized)
			return
		}
	})
}

// func loggingMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Do stuff here
// 		logger.DebugLogger.Println(r.RequestURI)
// 		logger.DebugLogger.Println(r.URL.RawQuery)
// 		// Call the next handler, which can be another middleware in the chain, or the final handler.
// 		next.ServeHTTP(w, r)
// 	})
// }
type RespJSON struct {
	Status int                    `json:"status"`
	Result map[string]interface{} `json:"result"`
}

func (respObj RespJSON) SetResponse() ([]byte, error) {
	jsonResponse, err := json.Marshal(respObj)
	if err != nil {
		fmt.Println("Unable to encode JSON")
	}
	return jsonResponse, err
}

func ReturnResponse(w http.ResponseWriter, resp []byte, httpStatus int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write(resp)
}
