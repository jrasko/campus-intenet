package api

import (
	"backend/model"
	"encoding/json"
	"net/http"
)

func (h Handler) PostNetworkHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var netConfig model.NetConfig
		err := json.NewDecoder(req.Body).Decode(&netConfig)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = h.networkService.CreateOrUpdateNetConfig(req.Context(), netConfig)
		if err != nil {
			sendHttpError(w, err)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func (h Handler) ListNetworkHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := model.NetworkRequestParams{
			Servers:  boolFilter(req, "server"),
			Disabled: boolFilter(req, "disabled"),
		}
		rooms, err := h.networkService.ListNetConfigs(req.Context(), params)
		if err != nil {
			sendHttpError(w, err)
			return
		}
		sendJSONResponse(w, rooms)
	}
}

func (h Handler) GetNetworkHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id, err := readIntFromVar(req, "id")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		rooms, err := h.networkService.GetNetConfig(req.Context(), id)
		if err != nil {
			sendHttpError(w, err)
			return
		}
		sendJSONResponse(w, rooms)
	}
}

func (h Handler) PutNetworkHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var netConfig model.NetConfig
		err := json.NewDecoder(req.Body).Decode(&netConfig)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		netConfig.ID, err = readIntFromVar(req, "id")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = h.networkService.CreateOrUpdateNetConfig(req.Context(), netConfig)
		if err != nil {
			sendHttpError(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (h Handler) DeleteNetworkHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id, err := readIntFromVar(req, "id")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = h.networkService.DeleteNetConfig(req.Context(), id)
		if err != nil {
			sendHttpError(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (h Handler) WriteConfigHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h.networkService.UpdateDhcpFile(r.Context())
		if err != nil {
			sendHttpError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h Handler) ToggleActivation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := readIntFromVar(r, "id")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = h.networkService.ToggleNetwork(r.Context(), id)
		if err != nil {
			sendHttpError(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
