package api

import (
	"backend/model"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type DhcpdService interface {
	UpdateConfig(ctx context.Context, config model.NetworkConfig) error
}

type DhcpdRepository interface {
	UpdateNetworkConfig(ctx context.Context, conf model.NetworkConfig) error
}

func NewRouter(
	service DhcpdService,
) http.Handler {
	router := mux.NewRouter()

	router.Handle("/dhcpd", PutConfigHandler(service)).
		Methods(http.MethodPut)

	return router
}
