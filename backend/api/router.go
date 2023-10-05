package api

import (
	"backend/model"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type DhcpService interface {
	UpdateMember(ctx context.Context, member model.MemberConfig) (model.MemberConfig, error)
	GetAllMembers(ctx context.Context, params model.RequestParams) ([]model.MemberConfig, error)
	GetMember(ctx context.Context, id int) (model.MemberConfig, error)
	DeleteMember(ctx context.Context, id int) error
	ResetPayment(ctx context.Context) error
	UpdateDhcpFile(ctx context.Context) error
	IsInconsistent() bool
	GetNotPayingMembers(ctx context.Context) ([]model.ReducedMember, error)
}

func NewRouter(config model.Configuration, service DhcpService) http.Handler {
	router := mux.NewRouter()

	auth := AuthHandler{config: config}
	h := Handler{service: service}

	router.
		Handle("/dhcp/login", auth.Login()).
		Methods(http.MethodPost)

	router.
		Handle("/dhcp", auth.Middleware(h.GetAllConfigHandler())).
		Methods(http.MethodGet)

	router.
		Handle("/dhcp", auth.Middleware(h.PostConfigHandler())).
		Methods(http.MethodPost)

	router.
		Handle("/dhcp/write", auth.Middleware(h.WriteConfigHandler())).
		Methods(http.MethodPost)

	router.
		Handle("/dhcp/resetPayment", auth.Middleware(h.ResetPaymentConfigHandler())).
		Methods(http.MethodPost)

	router.
		Handle("/dhcp/shame", auth.Middleware(h.WallOfShame())).
		Methods(http.MethodGet)

	router.
		Handle("/dhcp/{id}", auth.Middleware(h.GetConfigHandler())).
		Methods(http.MethodGet)

	router.
		Handle("/dhcp/{id}", auth.Middleware(h.DeleteConfigHandler())).
		Methods(http.MethodDelete)

	router.
		Handle("/dhcp/{id}", auth.Middleware(h.PutConfigHandler())).
		Methods(http.MethodPut)
	return router
}
