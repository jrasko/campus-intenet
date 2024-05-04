package api

import (
	"backend/model"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type MemberService interface {
	CreateOrUpdateMember(ctx context.Context, member model.Member) (model.Member, error)
	ListMembers(ctx context.Context, params model.MemberRequestParams) ([]model.Member, error)
	GetMember(ctx context.Context, id int) (model.Member, error)
	DeleteMember(ctx context.Context, id int) error

	GetNotPayingMembers(ctx context.Context) ([]model.ReducedMember, error)
	TogglePayment(ctx context.Context, id int) error
	ResetPayment(ctx context.Context) error
}
type RoomService interface {
	ListRooms(ctx context.Context, params model.RoomRequestParams) ([]model.Room, error)
}

type NetworkService interface {
	CreateOrUpdateNetConfig(ctx context.Context, config model.NetConfig) error
	ListNetConfigs(ctx context.Context, params model.NetworkRequestParams) ([]model.NetConfig, error)
	GetNetConfig(ctx context.Context, id int) (model.NetConfig, error)
	DeleteNetConfig(ctx context.Context, id int) error

	UpdateDhcpFile(ctx context.Context) error
	IsInconsistent() bool
}

func NewRouter(config model.Configuration, memberService MemberService, roomService RoomService, netService NetworkService) http.Handler {
	router := mux.NewRouter()

	auth := AuthHandler{config: config}
	h := Handler{
		memberService:  memberService,
		roomService:    roomService,
		networkService: netService,
	}

	router.
		Handle("/api/login", auth.Login()).
		Methods(http.MethodPost)

	router.
		Handle("/api/members", auth.Middleware(h.GetAllConfigHandler(), PermissionView)).
		Methods(http.MethodGet)
	router.
		Handle("/api/members", auth.Middleware(h.PostConfigHandler(), PermissionModify)).
		Methods(http.MethodPost)
	router.
		Handle("/api/members/{id}", auth.Middleware(h.GetConfigHandler(), PermissionView)).
		Methods(http.MethodGet)
	router.
		Handle("/api/members/{id}", auth.Middleware(h.DeleteConfigHandler(), PermissionModify)).
		Methods(http.MethodDelete)
	router.
		Handle("/api/members/{id}", auth.Middleware(h.PutConfigHandler(), PermissionModify)).
		Methods(http.MethodPut)

	router.
		Handle("/api/members/resetPayment", auth.Middleware(h.ResetPaymentConfigHandler(), PermissionFinance)).
		Methods(http.MethodPost)
	router.
		Handle("/api/members/{id}/togglePayment", auth.Middleware(h.TogglePayment(), PermissionFinance)).
		Methods(http.MethodPost)

	router.
		Handle("/api/shame", auth.Middleware(h.WallOfShame(), PermissionView)).
		Methods(http.MethodGet)

	router.
		Handle("/api/rooms", auth.Middleware(h.ListRooms(), PermissionView)).
		Methods(http.MethodGet)

	router.
		Handle("/api/write", auth.Middleware(h.WriteConfigHandler(), PermissionAdmin)).
		Methods(http.MethodPost)
	router.
		Handle("/api/net-configs", auth.Middleware(h.PostNetworkHandler(), PermissionAdmin)).
		Methods(http.MethodPost)
	router.
		Handle("/api/net-configs", auth.Middleware(h.ListNetworkHandler(), PermissionAdmin)).
		Methods(http.MethodGet)
	router.
		Handle("/api/net-configs/{id}", auth.Middleware(h.GetNetworkHandler(), PermissionAdmin)).
		Methods(http.MethodGet)
	router.
		Handle("/api/net-configs/{id}", auth.Middleware(h.PutNetworkHandler(), PermissionAdmin)).
		Methods(http.MethodPut)
	router.
		Handle("/api/net-configs/{id}", auth.Middleware(h.DeleteNetworkHandler(), PermissionAdmin)).
		Methods(http.MethodDelete)
	return router
}

func PermissionAdmin(role model.Role) bool {
	return role == model.RoleAdmin
}

func PermissionModify(role model.Role) bool {
	return role == model.RoleAdmin || role == model.RoleEditor
}

func PermissionFinance(role model.Role) bool {
	return role == model.RoleAdmin || role == model.RoleFinance
}

func PermissionView(_ model.Role) bool {
	return true
}
