package service

import (
	"backend/model"
	"context"
	"fmt"
	"log"
)

func (s *Service) IsInconsistent() bool {
	return s.inconsistentState
}

func (s *Service) UpdateDhcpFile(ctx context.Context) error {
	disabled := false
	users, err := s.netRepo.ListNetConfigs(ctx, model.NetworkRequestParams{Disabled: &disabled})
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

func (s *Service) CreateOrUpdateNetConfig(ctx context.Context, config model.NetConfig) error {
	err := s.validate.Var(config.IP, model.IPValidation)
	if err != nil {
		return mapValidationError(err)
	}
	err = s.validate.Struct(config)
	if err != nil {
		return mapValidationError(err)
	}
	_, err = s.netRepo.CreateOrUpdateNetConfig(ctx, config)
	if err != nil {
		return err
	}
	return s.UpdateDhcpFile(ctx)
}

func (s *Service) ListNetConfigs(ctx context.Context, params model.NetworkRequestParams) ([]model.NetConfig, error) {
	return s.netRepo.ListNetConfigs(ctx, params)
}

func (s *Service) GetNetConfig(ctx context.Context, id int) (model.NetConfig, error) {
	return s.netRepo.GetNetConfig(ctx, id)
}

func (s *Service) DeleteNetConfig(ctx context.Context, id int) error {
	err := s.netRepo.DeleteNetConfig(ctx, id)
	if err != nil {
		return err
	}
	return s.UpdateDhcpFile(ctx)
}
