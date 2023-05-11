package service

import (
	"backend/model"
	"backend/repository"
	"backend/service/allocation"
	"context"
	"net"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	validate    *validator.Validate
	networkRepo ResidentsRepository
	ipService   IPAllocationService
}

type IPAllocationService interface {
	GetUnusedIP(ctx context.Context) (net.IP, error)
	UnAllocateIP(ctx context.Context, ip net.IP) error
}

type ResidentsRepository interface {
	UpdateNetworkConfig(ctx context.Context, conf model.NetworkConfig) (model.NetworkConfig, error)
	GetAllNetworkConfigs(ctx context.Context) ([]model.NetworkConfig, error)
	GetNetworkConfig(ctx context.Context, mac string) (model.NetworkConfig, error)
	DeleteNetworkConfig(ctx context.Context, mac string) error
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

func (s Service) UpdateConfig(ctx context.Context, config model.NetworkConfig) (model.NetworkConfig, error) {
	err := s.validate.Struct(config)
	if err != nil {
		return model.NetworkConfig{}, err
	}

	_, err = s.networkRepo.GetNetworkConfig(ctx, config.Mac)
	if err == repository.ErrNotFound {
		config.IP, err = s.ipService.GetUnusedIP(ctx)
	} else if err != nil {
		return model.NetworkConfig{}, err
	}

	if err != nil {
		return model.NetworkConfig{}, err
	}
	return s.networkRepo.UpdateNetworkConfig(ctx, specialize(config))
}

func (s Service) GetAllConfigs(ctx context.Context) ([]model.NetworkConfig, error) {
	return s.networkRepo.GetAllNetworkConfigs(ctx)
}

func (s Service) GetConfig(ctx context.Context, mac string) (model.NetworkConfig, error) {
	return s.networkRepo.GetNetworkConfig(ctx, mac)
}

func (s Service) DeleteConfig(ctx context.Context, mac string) error {
	config, err := s.networkRepo.GetNetworkConfig(ctx, mac)
	if err != nil {
		return err
	}

	err = s.ipService.UnAllocateIP(ctx, config.IP)
	if err != nil {
		return err
	}

	return s.networkRepo.DeleteNetworkConfig(ctx, mac)
}

func (s Service) ResetPayment(ctx context.Context) error {
	return s.networkRepo.ResetPayment(ctx)
}
