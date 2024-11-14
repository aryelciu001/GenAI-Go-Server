package main

import (
	"context"

	"cloud.google.com/go/firestore"
	"ex.com/basicws/internal/constants"
)

func initDb() *firestore.Client {
	dbClient, err := firestore.NewClientWithDatabase(
		context.Background(), constants.PROJECT_ID, constants.FIRESTORE_DB_ID,
	)
	if err != nil {
		panic(err.Error())
	}
	return dbClient
}
