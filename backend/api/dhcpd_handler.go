package api

import (
	"backend/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const StatusInconsistent = 210

type Handler struct {
	service DhcpService
}

func (h Handler) PostConfigHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var member model.MemberConfig
		err := json.NewDecoder(r.Body).Decode(&member)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		member, err = h.service.UpdateMember(r.Context(), member)
		if err != nil {
			sendHttpError(w, err)
			return
		}

		sendJSONResponse(w, member)
	}
}

func (h Handler) PutConfigHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam, ok := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idParam)
		if !ok || err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var member model.MemberConfig
		err = json.NewDecoder(r.Body).Decode(&member)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		member.ID = id
		member, err = h.service.UpdateMember(r.Context(), member)
		if err != nil {
			sendHttpError(w, err)
			return
		}

		sendJSONResponse(w, member)
	}
}

func (h Handler) GetAllConfigHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := model.RequestParams{
			Search: r.FormValue("search"),
			Order:  r.FormValue("order"),
		}
		members, err := h.service.GetAllMembers(r.Context(), params)
		if err != nil {
			sendHttpError(w, err)
			return
		}

		if h.service.IsInconsistent() {
			w.WriteHeader(StatusInconsistent)
		}
		sendJSONResponse(w, members)
	}
}

func (h Handler) GetConfigHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam, ok := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idParam)
		if !ok || err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		member, err := h.service.GetMember(r.Context(), id)
		if err != nil {
			sendHttpError(w, err)
			return
		}

		if h.service.IsInconsistent() {
			w.WriteHeader(StatusInconsistent)
		}
		sendJSONResponse(w, member)
	}
}

func (h Handler) DeleteConfigHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam, ok := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idParam)
		if !ok || err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = h.service.DeleteMember(r.Context(), id)
		if err != nil {
			sendHttpError(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func (h Handler) ResetPaymentConfigHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h.service.ResetPayment(r.Context())
		if err != nil {
			sendHttpError(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func (h Handler) WriteConfigHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h.service.UpdateDhcpFile(r.Context())
		if err != nil {
			sendHttpError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h Handler) WallOfShame() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		members, err := h.service.GetNotPayingMembers(r.Context())
		if err != nil {
			sendHttpError(w, err)
			return
		}
		sendJSONResponse(w, members)
	}
}
