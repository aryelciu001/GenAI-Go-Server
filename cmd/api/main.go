package main

import (
	"log"
	"net/http"

	"ex.com/basicws/internal/config"
	"ex.com/basicws/internal/service"
)

func main() {
	cfg := config.LoadConfig()

	dbService := service.MustInitDb(cfg)
	defer dbService.Client.Close()

	vertexAIService := service.MustInitializeVertexAIService(cfg)
	defer vertexAIService.GenAIClient.Close()

	cloudStorageService := service.MustInitCloudStorageClient()
	defer cloudStorageService.Client.Close()

	redisService := service.InitRedisService()
	defer redisService.Client.Close()

	router := initRouter(cfg, dbService, vertexAIService, cloudStorageService, redisService)

	log.Printf("Listening to port %s", cfg.ServerPort)
	http.ListenAndServe(cfg.ServerPort, router)
}
