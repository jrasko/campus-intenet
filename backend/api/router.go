package api

import (
	"backend/model"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type DhcpService interface {
	CreateOrUpdateMember(ctx context.Context, member model.Member) (model.Member, error)
	ListMembers(ctx context.Context, params model.MemberRequestParams) ([]model.Member, error)
	GetMember(ctx context.Context, id int) (model.Member, error)
	DeleteMember(ctx context.Context, id int) error
	ResetPayment(ctx context.Context) error
	UpdateDhcpFile(ctx context.Context) error
	IsInconsistent() bool
	TogglePayment(ctx context.Context, id int) error
	GetNotPayingMembers(ctx context.Context) ([]model.ReducedMember, error)

	ListRooms(ctx context.Context, params model.RoomRequestParams) ([]model.Room, error)
}

func NewRouter(config model.Configuration, service DhcpService) http.Handler {
	router := mux.NewRouter()

	auth := AuthHandler{config: config}
	h := Handler{service: service}

	router.
		Handle("/dhcp/login", auth.Login()).
		Methods(http.MethodPost)

	router.
		Handle("/dhcp", auth.Middleware(h.GetAllConfigHandler(), PermissionView)).
		Methods(http.MethodGet)

	router.
		Handle("/dhcp", auth.Middleware(h.PostConfigHandler(), PermissionModify)).
		Methods(http.MethodPost)

	router.
		Handle("/dhcp/write", auth.Middleware(h.WriteConfigHandler(), PermissionModify)).
		Methods(http.MethodPost)

	router.
		Handle("/dhcp/resetPayment", auth.Middleware(h.ResetPaymentConfigHandler(), PermissionFinance)).
		Methods(http.MethodPost)

	router.
		Handle("/dhcp/shame", auth.Middleware(h.WallOfShame(), PermissionView)).
		Methods(http.MethodGet)

	router.
		Handle("/dhcp/{id}", auth.Middleware(h.GetConfigHandler(), PermissionView)).
		Methods(http.MethodGet)

	router.
		Handle("/dhcp/{id}", auth.Middleware(h.DeleteConfigHandler(), PermissionModify)).
		Methods(http.MethodDelete)

	router.
		Handle("/dhcp/{id}", auth.Middleware(h.PutConfigHandler(), PermissionModify)).
		Methods(http.MethodPut)

	router.
		Handle("/dhcp/{id}/togglePayment", auth.Middleware(h.TogglePayment(), PermissionFinance)).
		Methods(http.MethodPost)

	router.
		Handle("/api/rooms", auth.Middleware(h.ListRooms(), PermissionView)).
		Methods(http.MethodGet)
	return router
}

func PermissionModify(role model.Role) bool {
	return role == model.RoleAdmin
}

func PermissionFinance(role model.Role) bool {
	return role == model.RoleAdmin || role == model.RoleFinance
}

func PermissionView(_ model.Role) bool {
	return true
}
