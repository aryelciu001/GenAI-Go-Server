package main

import (
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"ex.com/basicws/internal/constants"
	"ex.com/basicws/internal/service"
)

type app struct {
	Db                  *firestore.Client
	VertexAIService     *service.VertexAIService
	CloudStorageService *service.CloudStorageService
	RedisService        *service.RedisService
}

func main() {
	db := initDb()
	defer db.Close()

	vertexAIService := service.InitializeVertexAIService()
	defer vertexAIService.GenAIClient.Close()

	cloudStorageService := service.InitCloudStorageClient()
	defer cloudStorageService.Client.Close()

	redisService := service.InitRedisService()
	defer redisService.Client.Close()

	router := initRouter(db, vertexAIService, cloudStorageService, redisService)

	fmt.Printf("Listening to port %s", constants.SERVER_PORT)
	http.ListenAndServe(constants.SERVER_PORT, router)
}
