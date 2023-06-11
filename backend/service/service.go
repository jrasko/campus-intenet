package service

import (
	"backend/model"
	"backend/service/allocation"
	"backend/service/confwriter"
	"context"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	validate    *validator.Validate
	memberRepo  MemberRepository
	ipService   IPAllocationService
	dhcpdWriter ConfWriter

	inconsistentState bool // indicates if the db has another state than the dhcpd file
}

type ConfWriter interface {
	WhitelistMacs(macs []string) error
}

type IPAllocationService interface {
	GetUnusedIP(ctx context.Context) (string, error)
}

type MemberRepository interface {
	allocation.IPRepository
	UpdateMemberConfig(ctx context.Context, conf model.MemberConfig) (model.MemberConfig, error)
	GetAllMemberConfigs(ctx context.Context) ([]model.MemberConfig, error)
	GetMemberConfig(ctx context.Context, id int) (model.MemberConfig, error)
	DeleteMemberConfig(ctx context.Context, id int) error
	ResetPayment(ctx context.Context) error
	GetAllMacs(ctx context.Context) ([]string, error)
	GetNonPayingMembers(ctx context.Context) ([]model.MemberConfig, error)
}

func New(config model.Configuration, repo MemberRepository) Service {
	dhcpdWriter := confwriter.New()
	s := Service{
		inconsistentState: false,
		memberRepo:        repo,
		dhcpdWriter:       dhcpdWriter,
		validate:          validator.New(),
		ipService:         allocation.New(repo, config.CIDR),
	}

	// generate config from db initially
	err := s.UpdateDhcpdFile(context.Background())
	if err != nil {
		panic(err)
	}
	return s
}

func (s Service) UpdateMember(ctx context.Context, config model.MemberConfig) (model.MemberConfig, error) {
	err := s.validate.Struct(config)
	if err != nil {
		if fieldErrors, ok := err.(validator.ValidationErrors); ok {
			message := ""
			for _, fieldError := range fieldErrors {
				message += fmt.Sprintf("%s:%s; ", fieldError.Field(), fieldError.Tag())
			}
			return model.MemberConfig{}, model.Error(http.StatusBadRequest, err.Error(), fmt.Sprintf(message))
		}
		return model.MemberConfig{}, err
	}

	if config.IP == "" {
		config.IP, err = s.ipService.GetUnusedIP(ctx)
		if err != nil {
			return model.MemberConfig{}, err
		}
	}

	memberConfig, err := s.memberRepo.UpdateMemberConfig(ctx, specialize(config))
	if err != nil {
		return model.MemberConfig{}, err
	}

	err = s.UpdateDhcpdFile(ctx)
	if err != nil {
		return model.MemberConfig{}, err
	}

	return memberConfig, err
}

func (s Service) UpdateDhcpdFile(ctx context.Context) error {
	macs, err := s.memberRepo.GetAllMacs(ctx)
	if err != nil {
		return err
	}

	err = s.dhcpdWriter.WhitelistMacs(macs)
	s.inconsistentState = err != nil
	return err
}

func (s Service) GetAllMembers(ctx context.Context) ([]model.MemberConfig, error) {
	return s.memberRepo.GetAllMemberConfigs(ctx)
}

func (s Service) GetMember(ctx context.Context, id int) (model.MemberConfig, error) {
	return s.memberRepo.GetMemberConfig(ctx, id)
}

func (s Service) DeleteMember(ctx context.Context, id int) error {
	err := s.memberRepo.DeleteMemberConfig(ctx, id)
	if err != nil {
		return err
	}

	return s.UpdateDhcpdFile(ctx)
}

func (s Service) ResetPayment(ctx context.Context) error {
	return s.memberRepo.ResetPayment(ctx)
}

func (s Service) IsInconsistent() bool {
	return s.inconsistentState
}

func (s Service) GetNotPayingMembers(ctx context.Context) ([]model.ReducedMember, error) {
	bobs, err := s.memberRepo.GetNonPayingMembers(ctx)
	if err != nil {
		return nil, err
	}

	person := make([]model.ReducedMember, 0, len(bobs))
	for _, member := range bobs {
		person = append(person, member.ToReduced())
	}
	return person, nil
}
