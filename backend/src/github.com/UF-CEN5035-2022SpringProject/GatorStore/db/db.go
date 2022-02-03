package db

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/UF-CEN5035-2022SpringProject/GatorStore/logger"
	"google.golang.org/api/option"
)

var DatabaseClient *firestore.Client

func ConnectionSetUp() {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("../uf-cen5035-se-firebase-adminsdk-ziukh-6d15950729.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		logger.ErrorLogger.Fatalln(err)
		log.Fatalln(err)
	}

	DatabaseClient, err := app.Firestore(ctx)
	if err != nil {
		logger.ErrorLogger.Fatalln(err)
		log.Fatalln(err)
	}
	defer DatabaseClient.Close()
}
