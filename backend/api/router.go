package api

import (
	"backend/model"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type DhcpdService interface {
	UpdateMember(ctx context.Context, member model.MemberConfig) (model.MemberConfig, error)
	GetAllMembers(ctx context.Context) ([]model.MemberConfig, error)
	GetMember(ctx context.Context, id int) (model.MemberConfig, error)
	DeleteMember(ctx context.Context, id int) error
	ResetPayment(ctx context.Context) error
	UpdateDhcpdFile(ctx context.Context) error
	IsInconsistent() bool
}

type DhcpdRepository interface{}

func NewRouter(service DhcpdService, config model.Configuration) http.Handler {
	router := mux.NewRouter()

	auth := AuthHandler{config: config}
	h := Handler{service: service}

	router.
		Handle("/dhcpd/login", auth.Login()).
		Methods(http.MethodPost)

	router.
		Handle("/dhcpd", auth.Middleware(h.GetAllConfigHandler())).
		Methods(http.MethodGet)

	router.
		Handle("/dhcpd", auth.Middleware(h.PostConfigHandler())).
		Methods(http.MethodPost)

	router.
		Handle("/dhcpd/write", auth.Middleware(h.WriteConfigHandler())).
		Methods(http.MethodPost)

	router.
		Handle("/dhcpd/resetPayment", auth.Middleware(h.ResetPaymentConfigHandler())).
		Methods(http.MethodPost)

	router.
		Handle("/dhcpd/{id}", auth.Middleware(h.GetConfigHandler())).
		Methods(http.MethodGet)

	router.
		Handle("/dhcpd/{id}", auth.Middleware(h.DeleteConfigHandler())).
		Methods(http.MethodDelete)

	router.
		Handle("/dhcpd/{id}", auth.Middleware(h.PutConfigHandler())).
		Methods(http.MethodPut)
	return router
}
