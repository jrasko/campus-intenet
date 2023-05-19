package api

import (
	"backend/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func PostConfigHandler(service DhcpdService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var config model.MemberConfig
		err := json.NewDecoder(r.Body).Decode(&config)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		config, err = service.UpdateConfig(r.Context(), config)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendJSONResponse(w, config)
	}
}

func PutConfigHandler(service DhcpdService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam, ok := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idParam)
		if !ok || err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var config model.MemberConfig
		err = json.NewDecoder(r.Body).Decode(&config)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		config.ID = id
		config, err = service.UpdateConfig(r.Context(), config)
		if err != nil {
			fmt.Println(err)
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
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendJSONResponse(w, configs)
	}
}

func GetConfigHandler(service DhcpdService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam, ok := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idParam)
		if !ok || err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		config, err := service.GetConfig(r.Context(), id)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendJSONResponse(w, config)
	}
}

func DeleteConfigHandler(service DhcpdService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam, ok := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idParam)
		if !ok || err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = service.DeleteConfig(r.Context(), id)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(204)
	}
}

func ResetPaymentConfigHandler(service DhcpdService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := service.ResetPayment(r.Context())
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(204)
	}
}

func sendJSONResponse(w http.ResponseWriter, v any) {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
