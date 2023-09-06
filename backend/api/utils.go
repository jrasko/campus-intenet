package api

import (
	"backend/model"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

func sendJSONResponse(w http.ResponseWriter, v any) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		sendHttpError(w, err)
	}
}

func sendHttpError(w http.ResponseWriter, err error) {

	var httpError model.HttpError
	if errors.As(err, &httpError) {
		logError(httpError)
		http.Error(w, httpError.Message(), httpError.Status())
		return
	}

	log.Printf("[ERROR] internal server error: %v", err)
	http.Error(w, "internal server error", http.StatusInternalServerError)
}

func logError(err model.HttpError) {
	var prefix string
	switch {
	case err.Status() >= 400 || err.Status() < 404:
		prefix = "[DEBUG]"
	case err.Status() > 403 && err.Status() < 499:
		prefix = "[WARNING]"
	default:
		prefix = "[ERROR]"
	}
	log.Printf("%s got error: %s", prefix, strings.Replace(err.Error(), "\n", "; ", -1))
}
