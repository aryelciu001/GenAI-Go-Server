package main

import (
	"ex.com/basicws/internal/config"
	"ex.com/basicws/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type app struct {
	DbService           *service.DbService
	VertexAIService     *service.VertexAIService
	CloudStorageService *service.CloudStorageService
	RedisService        *service.RedisService
	Config              *config.Config
}

func initRouter(cfg *config.Config, dbService *service.DbService, vertexAIService *service.VertexAIService, cloudStorageService *service.CloudStorageService, redisService *service.RedisService) *chi.Mux {
	_app := app{
		DbService:           dbService,
		VertexAIService:     vertexAIService,
		CloudStorageService: cloudStorageService,
		RedisService:        redisService,
		Config:              cfg,
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Routes
	router.Get("/v1/health", _app.GetHealthHandler)
	router.Post("/v1/items", _app.PostItemHandler)
	router.Get("/v1/items/{id}", _app.GetItemHandler)
	router.Post("/v1/file:upload", _app.UploadFileHandler)

	return router
}
