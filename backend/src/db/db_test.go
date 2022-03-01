package db

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	firebase "firebase.google.com/go"
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

// check the use info retrieving
func TestRequestUserObj(t *testing.T) {
	// expected := "dummy user data"
	// svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, expected)
	// }))
	// defer svr.Close()
	// c := NewClient(svr.URL)
	// res, err := c.UpperCase("anything")
	// if err != nil {
	// 	t.Errorf("expected err to be nil got %v", err)
	// }
	// res = strings.TrimSpace(res)
	// if res != expected {
	// 	t.Errorf("expected res to be %s got %s", expected, res)
	// }
}
