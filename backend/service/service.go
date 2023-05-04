package service

import (
	"backend/model"
	"backend/repository"
	"context"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	validate    *validator.Validate
	networkRepo NetworkRepository
}

type NetworkRepository interface {
	UpdateNetworkConfig(ctx context.Context, conf model.NetworkConfig) (model.NetworkConfig, error)
	GetAllNetworkConfigs(ctx context.Context) ([]model.NetworkConfig, error)
}

func New(repo repository.NetworkRepository) Service {
	v := validator.New()
	return Service{
		validate:    v,
		networkRepo: repo,
	}
}

func (s Service) UpdateConfig(ctx context.Context, config model.NetworkConfig) (model.NetworkConfig, error) {
	err := s.validate.Struct(config)
	if err != nil {
		return model.NetworkConfig{}, err
	}
	return s.networkRepo.UpdateNetworkConfig(ctx, config)
}

func (s Service) GetAllConfigs(ctx context.Context) ([]model.NetworkConfig, error) {
	return s.networkRepo.GetAllNetworkConfigs(ctx)
}
