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
			}
			err = service.UpdateConfig(nil, config)
			if err != nil {
				http.Error(resp, err.Error(), http.StatusInternalServerError)
			}
		})

}
