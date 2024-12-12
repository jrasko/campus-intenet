package service

import (
	"backend/model"
	"backend/service/allocation"
	"backend/service/confwriter"
	"context"
	"log"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	validate   *validator.Validate
	memberRepo MemberRepository
	roomRepo   RoomRepository
	netRepo    NetworkRepository
	ipService  IPAllocationService
	dhcpWriter ConfWriter

	inconsistentState bool // indicates if the db has another state than the user-list.json file
}

type ConfWriter interface {
	WhitelistMacs(netConfigs []model.NetConfig) error
}

type IPAllocationService interface {
	GetUnusedIP(ctx context.Context) (string, error)
}

type MemberRepository interface {
	ResetPayment(ctx context.Context) error

	CreateOrUpdateMember(ctx context.Context, member model.Member) (model.Member, error)
	GetMember(ctx context.Context, id int) (model.Member, error)
	DeleteMember(ctx context.Context, id int) error
}

type RoomRepository interface {
	GetRoom(ctx context.Context, number string) (model.Room, error)
	ListRooms(ctx context.Context, params model.RoomRequestParams) ([]model.Room, error)
}

type NetworkRepository interface {
	CreateOrUpdateNetConfig(ctx context.Context, config model.NetConfig) (model.NetConfig, error)
	GetNetConfig(ctx context.Context, id int) (model.NetConfig, error)
	ListNetConfigs(ctx context.Context, params model.NetworkRequestParams) ([]model.NetConfig, error)
	DeleteNetConfig(ctx context.Context, id int) error
	Deactivate(ctx context.Context, ids []int) error
}

func New(
	memberRepo MemberRepository,
	roomRepo RoomRepository,
	netRepo NetworkRepository,
	jsonWriter confwriter.JsonWriter,
	ipAllocation allocation.Service,
) *Service {
	s := Service{
		validate:          validator.New(),
		memberRepo:        memberRepo,
		roomRepo:          roomRepo,
		netRepo:           netRepo,
		dhcpWriter:        jsonWriter,
		ipService:         ipAllocation,
		inconsistentState: false,
	}

	// generate config from db initially
	err := s.UpdateDhcpFile(context.Background())
	if err != nil {
		log.Printf("[ERROR] when updating dhcp file: %#v", err)
		panic(err)
	}
	return &s
}
