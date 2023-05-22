package api

import (
	"backend/model"
	"encoding/json"
	"log"
	"net/http"
)

func sendJSONResponse(w http.ResponseWriter, v any) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		sendHttpError(w, err)
	}
}

func sendHttpError(w http.ResponseWriter, err error) {
	log.Println(err)
	if httpError, ok := err.(model.HttpError); ok {
		http.Error(w, httpError.Message(), httpError.Status())
		return
	}

	http.Error(w, "internal server error", http.StatusInternalServerError)
}
