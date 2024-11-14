package main

import (
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"ex.com/basicws/internal/constants"
	"ex.com/basicws/internal/service"
)

type app struct {
	Db              *firestore.Client
	VertexAIService *service.VertexAIService
}

func main() {
	db := initDb()
	vertexAIService := service.InitializeVertexAIService()

	defer db.Close()
	defer vertexAIService.GenAIClient.Close()

	router := initRouter(db, &vertexAIService)

	fmt.Printf("Listening to port %s", constants.PORT)
	http.ListenAndServe(constants.PORT, router)
}
