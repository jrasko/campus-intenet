package service

import (
	"backend/model"
	"backend/repository"
	"backend/service/allocation"
	"context"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	validate    *validator.Validate
	networkRepo ResidentsRepository
	ipService   IPAllocationService
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
}

func New(repo repository.NetworkRepository) Service {
	v := validator.New()
	return Service{
		validate:    v,
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

	return s.networkRepo.UpdateNetworkConfig(ctx, specialize(config))
}

func (s Service) GetAllConfigs(ctx context.Context) ([]model.MemberConfig, error) {
	return s.networkRepo.GetAllNetworkConfigs(ctx)
}

func (s Service) GetConfig(ctx context.Context, id int) (model.MemberConfig, error) {
	return s.networkRepo.GetNetworkConfig(ctx, id)
}

func (s Service) DeleteConfig(ctx context.Context, id int) error {
	return s.networkRepo.DeleteNetworkConfig(ctx, id)
}

func (s Service) ResetPayment(ctx context.Context) error {
	return s.networkRepo.ResetPayment(ctx)
}
