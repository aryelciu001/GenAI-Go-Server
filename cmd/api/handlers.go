package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ex.com/basicws/internal/constants"
	"github.com/go-chi/chi/v5"
)

func (app *app) GetHealthHandler(w http.ResponseWriter, r *http.Request) {
	res, err := app.VertexAIService.GenerateText(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ret := map[string]string{
		"status": "ok",
		"poem":   res,
	}
	encodedRet, err := json.Marshal(ret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(encodedRet)
}

func (app *app) PostItemHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	reqBody := make(map[string]interface{})
	parsingError := json.NewDecoder(r.Body).Decode(&reqBody)

	if parsingError != nil {
		http.Error(w, fmt.Sprintf("Payload parsing error: %s", parsingError.Error()), http.StatusBadRequest)
		return
	}

	docRef, _, err := app.Db.Collection("poki").Add(ctx, reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	snapshot, err := docRef.Get(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	doc, err := json.Marshal(snapshot.Data())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(doc)
}

func (app *app) GetItemHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "No ID is passed", http.StatusBadRequest)
		return
	}

	snapshot, err := app.Db.Collection("poki").Doc(id).Get(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := snapshot.Data()
	encodedData, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(encodedData)
}

func (app *app) UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	parsingError := r.ParseMultipartForm(10 * 1000000)
	if parsingError != nil {
		http.Error(w, fmt.Sprintf("failed parsing body: %v", parsingError.Error()), http.StatusBadRequest)
		return
	}

	file, fileHandler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed parsing body: %v", err.Error()), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	err = app.CloudStorageService.UploadToBucket(ctx, constants.BUCKET_NAME, file, fileHandler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("OK"))
}
