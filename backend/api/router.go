package api

import (
	"backend/model"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type DhcpdService interface {
	UpdateConfig(ctx context.Context, config model.MemberConfig) (model.MemberConfig, error)
	GetAllConfigs(ctx context.Context) ([]model.MemberConfig, error)
	GetConfig(ctx context.Context, id int) (model.MemberConfig, error)
	DeleteConfig(ctx context.Context, id int) error
	ResetPayment(ctx context.Context) error
}

type DhcpdRepository interface{}

func NewRouter(
	service DhcpdService,
	config model.Configuration,
) http.Handler {
	router := mux.NewRouter()

	router.
		Handle("/dhcpd/login", Login(config)).
		Methods(http.MethodPost)

	router.
		Handle("/dhcpd", AuthMiddleware(PostConfigHandler(service), config)).
		Methods(http.MethodPost)

	router.
		Handle("/dhcpd/{id}", AuthMiddleware(PutConfigHandler(service), config)).
		Methods(http.MethodPut)

	router.
		Handle("/dhcpd", AuthMiddleware(GetAllConfigHandler(service), config)).
		Methods(http.MethodGet)

	router.
		Handle("/dhcpd/resetPayment", AuthMiddleware(ResetPaymentConfigHandler(service), config)).
		Methods(http.MethodPost)

	router.
		Handle("/dhcpd/{id}", AuthMiddleware(GetConfigHandler(service), config)).
		Methods(http.MethodGet)

	router.
		Handle("/dhcpd/{id}", AuthMiddleware(DeleteConfigHandler(service), config)).
		Methods(http.MethodDelete)

	return router
}
