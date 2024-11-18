package service

import (
	"context"

	"cloud.google.com/go/firestore"
	"ex.com/basicws/internal/constants"
)

type DbService struct {
	Client *firestore.Client
}

func MustInitDb() *DbService {
	dbClient, err := firestore.NewClientWithDatabase(
		context.Background(), constants.PROJECT_ID, constants.FIRESTORE_DB_ID,
	)
	if err != nil {
		panic(err.Error())
	}
	return &DbService{
		Client: dbClient,
	}
}
