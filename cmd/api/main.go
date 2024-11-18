package main

import (
	"fmt"
	"net/http"

	"ex.com/basicws/internal/constants"
	"ex.com/basicws/internal/service"
)

type app struct {
	DbService           *service.DbService
	VertexAIService     *service.VertexAIService
	CloudStorageService *service.CloudStorageService
	RedisService        *service.RedisService
}

func main() {
	dbService := service.MustInitDb()
	defer dbService.Client.Close()

	vertexAIService := service.MustInitializeVertexAIService()
	defer vertexAIService.GenAIClient.Close()

	cloudStorageService := service.MustInitCloudStorageClient()
	defer cloudStorageService.Client.Close()

	redisService := service.InitRedisService()
	defer redisService.Client.Close()

	router := initRouter(dbService, vertexAIService, cloudStorageService, redisService)

	fmt.Printf("Listening to port %s", constants.SERVER_PORT)
	http.ListenAndServe(constants.SERVER_PORT, router)
}
