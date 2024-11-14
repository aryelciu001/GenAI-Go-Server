package main

import (
	"context"

	"cloud.google.com/go/firestore"
)

func initDb() *firestore.Client {
	dbClient, err := firestore.NewClientWithDatabase(
		context.Background(), PROJECT_ID, FIRESTORE_DB_ID,
	)
	if err != nil {
		panic(err.Error())
	}
	return dbClient
}
