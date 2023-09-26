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
	WhitelistUsers(macs []model.MemberConfig) error
}

type IPAllocationService interface {
	GetUnusedIP(ctx context.Context) (string, error)
}

type MemberRepository interface {
	allocation.IPRepository
	UpdateMemberConfig(ctx context.Context, conf model.MemberConfig) (model.MemberConfig, error)
	GetAllMemberConfigs(ctx context.Context, params model.RequestParams) ([]model.MemberConfig, error)
	GetMemberConfig(ctx context.Context, id int) (model.MemberConfig, error)
	DeleteMemberConfig(ctx context.Context, id int) error
	ResetPayment(ctx context.Context) error
	GetEnabledUsers(ctx context.Context) ([]model.MemberConfig, error)
	GetNonPayingMembers(ctx context.Context) ([]model.MemberConfig, error)
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

func (s *Service) UpdateMember(ctx context.Context, member model.MemberConfig) (model.MemberConfig, error) {
	err := s.validate.Struct(member)
	if err != nil {
		return model.MemberConfig{}, mapValidationError(err)
	}

	member.Sanitize()
	member.Manufacturer = ouiMappings[member.Mac[:8]]

	if member.IP == "" {
		member.IP, err = s.ipService.GetUnusedIP(ctx)
		if err != nil {
			return model.MemberConfig{}, err
		}
	}

	member, err = s.memberRepo.UpdateMemberConfig(ctx, specialize(member))
	if err != nil {
		return model.MemberConfig{}, model.WrapGormError(err)
	}

	err = s.UpdateDhcpFile(ctx)
	if err != nil {
		return model.MemberConfig{}, err
	}

	return member, err
}

func (s *Service) UpdateDhcpFile(ctx context.Context) error {
	users, err := s.memberRepo.GetEnabledUsers(ctx)
	if err != nil {
		return model.WrapGormError(err)
	}

	err = s.dhcpWriter.WhitelistUsers(users)
	if err != nil {
		s.inconsistentState = true
		return fmt.Errorf("when updating dhcp file: %w", err)
	}

	log.Println("[DEBUG] Successfully updated whitelist")
	s.inconsistentState = false
	return nil
}

func (s *Service) GetAllMembers(ctx context.Context, params model.RequestParams) ([]model.MemberConfig, error) {
	err := s.validate.Struct(params)
	if err != nil {
		return []model.MemberConfig{}, mapValidationError(err)
	}
	members, err := s.memberRepo.GetAllMemberConfigs(ctx, params)
	if err != nil {
		return []model.MemberConfig{}, model.WrapGormError(err)
	}
	return members, nil
}

func (s *Service) GetMember(ctx context.Context, id int) (model.MemberConfig, error) {
	member, err := s.memberRepo.GetMemberConfig(ctx, id)
	if err != nil {
		return model.MemberConfig{}, model.WrapGormError(err)
	}
	return member, nil
}

func (s *Service) DeleteMember(ctx context.Context, id int) error {
	err := s.memberRepo.DeleteMemberConfig(ctx, id)
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

func (s *Service) GetNotPayingMembers(ctx context.Context) ([]model.ReducedMember, error) {
	idiots, err := s.memberRepo.GetNonPayingMembers(ctx)
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
