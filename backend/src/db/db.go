package db

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
	"google.golang.org/api/option"
)

var FireBaseClient *firestore.Client
var DatabaseCtx context.Context

// Database Collections(Tables)
var DbCollections = map[string]string{
	"users":       "users",
	"stores":      "stores",
	"products":    "products",
	"settings":    "settings",
	"jwtTokenMap": "jwtTokenMap",
}

func ConnectionCreate() {
	// Use a service account
	DatabaseCtx = context.Background()
	sa := option.WithCredentialsFile("./db_secret.json")
	app, err := firebase.NewApp(DatabaseCtx, nil, sa)
	if err != nil {
		logger.ErrorLogger.Fatalln(err)
		log.Fatalln(err)
	}

	client, err := app.Firestore(DatabaseCtx)
	if err != nil {
		logger.ErrorLogger.Fatalln(err)
	}

	FireBaseClient = client
	logger.DebugLogger.Println("Successfully connect to FireStore")
}
