package service

import (
	"backend/model"
	"backend/repository"
	"backend/service/allocation"
	"backend/service/confwriter"
	"context"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	validate    *validator.Validate
	networkRepo ResidentsRepository
	ipService   IPAllocationService
	dhcpdWriter ConfWriter
}

type ConfWriter interface {
	WhitelistMacs(macs []string) error
}

type IPAllocationService interface {
	GetUnusedIP(ctx context.Context) (string, error)
}

type ResidentsRepository interface {
	UpdateNetworkConfig(ctx context.Context, conf model.MemberConfig) (model.MemberConfig, error)
	GetAllNetworkConfigs(ctx context.Context) ([]model.MemberConfig, error)
	GetNetworkConfig(ctx context.Context, id int) (model.MemberConfig, error)
	DeleteNetworkConfig(ctx context.Context, id int) error
	ResetPayment(ctx context.Context) error
	GetAllMacs(ctx context.Context) ([]string, error)
}

func New(repo repository.NetworkRepository) Service {
	w := confwriter.New()
	v := validator.New()
	return Service{
		validate:    v,
		dhcpdWriter: w,
		networkRepo: repo,
		ipService:   allocation.New(repo),
	}
}

func (s Service) UpdateConfig(ctx context.Context, config model.MemberConfig) (model.MemberConfig, error) {
	err := s.validate.Struct(config)
	if err != nil {
		return model.MemberConfig{}, err
	}

	if config.IP == "" {
		config.IP, err = s.ipService.GetUnusedIP(ctx)
		if err != nil {
			return model.MemberConfig{}, err
		}
	}

	//FIXME konsistenz?
	networkConfig, err := s.networkRepo.UpdateNetworkConfig(ctx, specialize(config))
	if err != nil {
		return model.MemberConfig{}, err
	}

	macs, err := s.networkRepo.GetAllMacs(ctx)
	if err != nil {
		return model.MemberConfig{}, err
	}

	err = s.dhcpdWriter.WhitelistMacs(macs)
	if err != nil {
		return model.MemberConfig{}, err
	}

	return networkConfig, err
}

func (s Service) GetAllConfigs(ctx context.Context) ([]model.MemberConfig, error) {
	return s.networkRepo.GetAllNetworkConfigs(ctx)
}

func (s Service) GetConfig(ctx context.Context, id int) (model.MemberConfig, error) {
	return s.networkRepo.GetNetworkConfig(ctx, id)
}

func (s Service) DeleteConfig(ctx context.Context, id int) error {
	err := s.networkRepo.DeleteNetworkConfig(ctx, id)
	if err != nil {
		return err
	}

	macs, err := s.networkRepo.GetAllMacs(ctx)
	if err != nil {
		return err
	}

	return s.dhcpdWriter.WhitelistMacs(macs)
}

func (s Service) ResetPayment(ctx context.Context) error {
	return s.networkRepo.ResetPayment(ctx)
}
