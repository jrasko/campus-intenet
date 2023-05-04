package api

import (
	"backend/model"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type DhcpdService interface {
	UpdateConfig(ctx context.Context, config model.NetworkConfig) (model.NetworkConfig, error)
	GetAllConfigs(ctx context.Context) ([]model.NetworkConfig, error)
}

type DhcpdRepository interface{}

func NewRouter(
	service DhcpdService,
) http.Handler {
	router := mux.NewRouter()

	router.
		Handle("/dhcpd", PutConfigHandler(service)).
		Methods(http.MethodPut)
	router.
		Handle("/dhcpd", GetAllConfigHandler(service)).
		Methods(http.MethodGet)

	return router
}
