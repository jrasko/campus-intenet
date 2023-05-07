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
	GetConfig(ctx context.Context, mac string) (model.NetworkConfig, error)
	DeleteConfig(ctx context.Context, mac string) error
}

type DhcpdRepository interface{}

func NewRouter(
	service DhcpdService,
) http.Handler {
	router := mux.NewRouter()

	router.
		Handle("/dhcpd", AuthMiddleware(PutConfigHandler(service))).
		Methods(http.MethodPut)
	router.
		Handle("/dhcpd", AuthMiddleware(GetAllConfigHandler(service))).
		Methods(http.MethodGet)
	router.
		Handle("/dhcpd/{mac}", AuthMiddleware(GetConfigHandler(service))).
		Methods(http.MethodGet)
	router.
		Handle("/dhcpd/{mac}", AuthMiddleware(DeleteConfigHandler(service))).
		Methods(http.MethodDelete)

	return router
}
