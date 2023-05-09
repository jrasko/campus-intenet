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

type Configuration struct {
	Username   string
	Password   string
	HMACSecret string
	Salt       string
}

type DhcpdRepository interface{}

func NewRouter(
	service DhcpdService,
	config Configuration,
) http.Handler {
	router := mux.NewRouter()

	router.
		Handle("/dhcpd/login", Login(config)).
		Methods(http.MethodPost)

	router.
		Handle("/dhcpd", AuthMiddleware(PutConfigHandler(service), config)).
		Methods(http.MethodPut)

	router.
		Handle("/dhcpd", AuthMiddleware(GetAllConfigHandler(service), config)).
		Methods(http.MethodGet)

	router.
		Handle("/dhcpd/resetPayment", AuthMiddleware(ResetPaymentConfigHandler(service), config)).
		Methods(http.MethodPost)

	router.
		Handle("/dhcpd/{mac}", AuthMiddleware(GetConfigHandler(service), config)).
		Methods(http.MethodGet)

	router.
		Handle("/dhcpd/{mac}", AuthMiddleware(DeleteConfigHandler(service), config)).
		Methods(http.MethodDelete)

	return router
}
