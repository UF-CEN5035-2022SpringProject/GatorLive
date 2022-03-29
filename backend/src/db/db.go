package db

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var FireBaseClient *firestore.Client
var DatabaseCtx context.Context

// Database Collections(Tables)
var DbCollections = map[string]string{
	"users":       "users",
	"stores":      "stores",
	"products":    "products",
	"orders":      "orders",
	"settings":    "settings",
	"jwtTokenMap": "jwtTokenMap",
	"lives":       "lives",
}

func ConnectionCreate() {
	cwd, _ := os.Getwd()
	DatabaseCtx = context.Background()
	sa := option.WithCredentialsFile(cwd + "/db_secret.json")
	app, err := firebase.NewApp(DatabaseCtx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(DatabaseCtx)
	if err != nil {
		log.Fatalln(err)
	}

	FireBaseClient = client
}
