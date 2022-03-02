package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	firebase "firebase.google.com/go"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/api"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/test"
	"google.golang.org/api/option"
)

// check the connection function
func TestDbConnection(t *testing.T) {
	path, _ := os.Getwd()
	fmt.Println(path)

	if _, err := os.Stat("../db_secret.json"); errors.Is(err, os.ErrNotExist) {
		t.Errorf("Db credential file missing, Error - %v", err)
	}

	DatabaseCtx = context.Background()
	sa := option.WithCredentialsFile("../db_secret.json")
	app, err := firebase.NewApp(DatabaseCtx, nil, sa)

	if err != nil {
		t.Errorf("Db connection initialize error, Error - %v", err)
	}

	_, err = app.Firestore(DatabaseCtx)
	if err != nil {
		t.Errorf("Db connect, Error - %v", err)
	}
}

func TestRequestUserObj(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/test/api/user/info", nil)
	w := httptest.NewRecorder()

	test.TestDBGetUserObj(w, r)
	resp := w.Result()
	defer resp.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	var response api.RespJSON
	err = json.Unmarshal(b, &response)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	fmt.Printf("resp %v", response)
}
