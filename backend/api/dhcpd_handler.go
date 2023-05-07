package api

import (
	"backend/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func PutConfigHandler(service DhcpdService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var config model.NetworkConfig
		err := json.NewDecoder(r.Body).Decode(&config)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		config, err = service.UpdateConfig(nil, config)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendJSONResponse(w, config)
	}
}

func GetAllConfigHandler(service DhcpdService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		configs, err := service.GetAllConfigs(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendJSONResponse(w, configs)
	}
}

func GetConfigHandler(service DhcpdService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mac, ok := mux.Vars(r)["mac"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		config, err := service.GetConfig(r.Context(), mac)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendJSONResponse(w, config)
	}
}

func DeleteConfigHandler(service DhcpdService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mac, ok := mux.Vars(r)["mac"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err := service.DeleteConfig(r.Context(), mac)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(200)
	}
}

func ResetPaymentConfigHandler(service DhcpdService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := service.ResetPayment(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(200)
	}
}

func sendJSONResponse(w http.ResponseWriter, v any) {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
