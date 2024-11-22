package service

import (
	"context"

	"cloud.google.com/go/firestore"
	"ex.com/basicws/internal/config"
)

type DbService struct {
	Client *firestore.Client
}

func MustInitDb(cfg *config.Config) *DbService {
	dbClient, err := firestore.NewClientWithDatabase(
		context.Background(), cfg.ProjectID, cfg.FirestoreDbID,
	)
	if err != nil {
		panic(err.Error())
	}
	return &DbService{
		Client: dbClient,
	}
}
