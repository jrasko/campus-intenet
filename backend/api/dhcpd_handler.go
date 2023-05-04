package api

import (
	"backend/model"
	"encoding/json"
	"net/http"
)

func PutConfigHandler(service DhcpdService) http.Handler {
	return http.HandlerFunc(
		func(resp http.ResponseWriter, req *http.Request) {
			var config model.NetworkConfig
			err := json.NewDecoder(req.Body).Decode(&config)
			if err != nil {
				http.Error(resp, err.Error(), http.StatusBadRequest)
				return
			}
			config, err = service.UpdateConfig(nil, config)
			if err != nil {
				http.Error(resp, err.Error(), http.StatusInternalServerError)
				return
			}
			err = json.NewEncoder(resp).Encode(config)
			if err != nil {
				http.Error(resp, err.Error(), http.StatusInternalServerError)
				return
			}
		})
}

func GetAllConfigHandler(service DhcpdService) http.Handler {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			configs, err := service.GetAllConfigs(request.Context())
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
			err = json.NewEncoder(writer).Encode(configs)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
		})
}
