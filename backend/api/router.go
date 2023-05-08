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
	ResetPayment(ctx context.Context) error
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

	router.
		Handle("/dhcpd/resetPayment", ResetPaymentConfigHandler(service)).
		Methods(http.MethodPost)

	router.
		Handle("/dhcpd/{mac}", GetConfigHandler(service)).
		Methods(http.MethodGet)

	router.
		Handle("/dhcpd/{mac}", DeleteConfigHandler(service)).
		Methods(http.MethodDelete)

	return router
}
