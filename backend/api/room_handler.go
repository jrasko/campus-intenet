package api

import (
	"backend/model"
	"net/http"
)

func (h Handler) ListRooms() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := model.RoomRequestParams{
			IsOccupied: boolFilter(req, "occupied"),
			Blocks:     req.URL.Query()["block"],
		}
		rooms, err := h.roomService.ListRooms(req.Context(), params)
		if err != nil {
			sendHttpError(w, err)
			return
		}
		sendJSONResponse(w, rooms)
	}
}
