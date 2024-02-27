package api

import (
	"backend/model"
	"encoding/json"
	"errors"
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
		var reqMember requestMember
		err := json.NewDecoder(r.Body).Decode(&reqMember)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		member := reqMember.toModel()
		member, err = h.service.CreateOrUpdateMember(r.Context(), member)
		if err != nil {
			sendHttpError(w, err)
			return
		}

		sendJSONResponse(w, member)
	}
}

func (h Handler) PutConfigHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var reqMember requestMember
		err := json.NewDecoder(r.Body).Decode(&reqMember)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		member := reqMember.toModel()
		member.ID, err = readIDFromVar(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		member, err = h.service.CreateOrUpdateMember(r.Context(), member)
		if err != nil {
			sendHttpError(w, err)
			return
		}

		sendJSONResponse(w, member)
	}
}

func (h Handler) GetAllConfigHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := model.MemberRequestParams{
			Search:   r.FormValue("search"),
			Disabled: boolFilter(r, "disabled"),
			HasPaid:  boolFilter(r, "hasPaid"),
		}
		members, err := h.service.ListMembers(r.Context(), params)
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
		id, err := readIDFromVar(r)
		if err != nil {
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
		id, err := readIDFromVar(r)
		if err != nil {
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

func (h Handler) TogglePayment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := readIDFromVar(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = h.service.TogglePayment(r.Context(), id)
		if err != nil {
			sendHttpError(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func readIDFromVar(r *http.Request) (int, error) {
	idParam, ok := mux.Vars(r)["id"]
	if !ok {
		return 0, errors.New("could not read id param")
	}
	return strconv.Atoi(idParam)
}

func boolFilter(r *http.Request, name string) *bool {
	value := r.FormValue(name)
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return nil
	}
	return &boolValue
}
