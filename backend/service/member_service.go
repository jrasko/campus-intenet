package service

import (
	"backend/model"
	"backend/service/allocation"
	"backend/service/confwriter"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	validate   *validator.Validate
	memberRepo MemberRepository
	ipService  IPAllocationService
	dhcpWriter ConfWriter

	inconsistentState bool // indicates if the db has another state than the user-list.json file
}

type ConfWriter interface {
	WhitelistMacs(netConfigs []model.DhcpConfig) error
}

type IPAllocationService interface {
	GetUnusedIP(ctx context.Context) (string, error)
}

type MemberRepository interface {
	CreateOrUpdateMember(ctx context.Context, member model.Member) (model.Member, error)
	GetMember(ctx context.Context, id int) (model.Member, error)
	ListMembers(ctx context.Context, params model.MemberRequestParams) ([]model.Member, error)
	DeleteMembers(ctx context.Context, id int) error
	ResetPayment(ctx context.Context) error

	GetEnabledNets(ctx context.Context) ([]model.DhcpConfig, error)
	ListRooms(ctx context.Context, params model.RoomRequestParams) ([]model.Room, error)
}

func New(repo MemberRepository, jsonWriter confwriter.JsonWriter, ipAllocation allocation.Service) *Service {
	s := Service{
		validate:          validator.New(),
		memberRepo:        repo,
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

func (s *Service) CreateOrUpdateMember(ctx context.Context, member model.Member) (model.Member, error) {
	err := s.validate.Struct(member)
	if err != nil {
		return model.Member{}, mapValidationError(err)
	}

	member.Sanitize()
	member.DhcpConfig.Manufacturer = ouiMappings[member.DhcpConfig.Mac[:8]]

	if member.DhcpConfig.IP == "" {
		member.DhcpConfig.IP, err = s.ipService.GetUnusedIP(ctx)
		if err != nil {
			return model.Member{}, err
		}
	}

	member.LastEditor, _ = ctx.Value(model.FieldUsername).(string)
	member, err = s.memberRepo.CreateOrUpdateMember(ctx, specialize(member))
	if err != nil {
		return model.Member{}, model.WrapGormError(err)
	}

	err = s.UpdateDhcpFile(ctx)
	if err != nil {
		return model.Member{}, err
	}

	return member, err
}

func (s *Service) UpdateDhcpFile(ctx context.Context) error {
	users, err := s.memberRepo.GetEnabledNets(ctx)
	if err != nil {
		return model.WrapGormError(err)
	}

	err = s.dhcpWriter.WhitelistMacs(users)
	if err != nil {
		s.inconsistentState = true
		return fmt.Errorf("when updating dhcp file: %w", err)
	}

	log.Println("[DEBUG] Successfully updated whitelist")
	s.inconsistentState = false
	return nil
}

func (s *Service) ListMembers(ctx context.Context, params model.MemberRequestParams) ([]model.Member, error) {
	err := s.validate.Struct(params)
	if err != nil {
		return []model.Member{}, mapValidationError(err)
	}
	members, err := s.memberRepo.ListMembers(ctx, params)
	if err != nil {
		return []model.Member{}, model.WrapGormError(err)
	}
	return members, nil
}

func (s *Service) GetMember(ctx context.Context, id int) (model.Member, error) {
	member, err := s.memberRepo.GetMember(ctx, id)
	if err != nil {
		return model.Member{}, model.WrapGormError(err)
	}
	return member, nil
}

func (s *Service) DeleteMember(ctx context.Context, id int) error {
	err := s.memberRepo.DeleteMembers(ctx, id)
	if err != nil {
		return model.WrapGormError(err)
	}

	return s.UpdateDhcpFile(ctx)
}

func (s *Service) ResetPayment(ctx context.Context) error {
	return model.WrapGormError(s.memberRepo.ResetPayment(ctx))
}

func (s *Service) IsInconsistent() bool {
	return s.inconsistentState
}

func (s *Service) TogglePayment(ctx context.Context, id int) error {
	config, err := s.memberRepo.GetMember(ctx, id)
	if err != nil {
		return err
	}
	config.HasPaid = !config.HasPaid
	_, err = s.memberRepo.CreateOrUpdateMember(ctx, config)
	return err
}

func (s *Service) GetNotPayingMembers(ctx context.Context) ([]model.ReducedMember, error) {
	hasPaid := false
	idiots, err := s.memberRepo.ListMembers(ctx, model.MemberRequestParams{HasPaid: &hasPaid})
	if err != nil {
		return nil, model.WrapGormError(err)
	}

	reducedIdiots := make([]model.ReducedMember, 0, len(idiots))
	for _, member := range idiots {
		reducedIdiots = append(reducedIdiots, member.ToReduced())
	}
	return reducedIdiots, nil
}

func mapValidationError(err error) error {
	var fieldErrors validator.ValidationErrors
	if errors.As(err, &fieldErrors) {
		message := ""
		for _, fieldError := range fieldErrors {
			message += fmt.Sprintf("%s:%s-%s; ", fieldError.Field(), fieldError.Tag(), fieldError.Param())
		}
		return model.Error(http.StatusBadRequest, err.Error(), message)
	}
	return err
}
