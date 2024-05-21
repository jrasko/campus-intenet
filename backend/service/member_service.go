package service

import (
	"backend/model"
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func (s *Service) CreateOrUpdateMember(ctx context.Context, member model.Member) (model.Member, error) {
	err := s.validate.Struct(member)
	if err != nil {
		return model.Member{}, mapValidationError(err)
	}

	// test if room exists
	_, err = s.roomRepo.GetRoom(ctx, member.RoomNr)
	if err != nil {
		return model.Member{}, model.WrapGormError(err)
	}

	member.Sanitize()
	member.LastEditor, _ = ctx.Value(model.FieldUsername).(string)
	member.NetConfig.Manufacturer = ouiMappings[member.NetConfig.Mac[:8]]

	if member.NetConfig.IP == "" {
		member.NetConfig.IP, err = s.ipService.GetUnusedIP(ctx)
		if err != nil {
			return model.Member{}, err
		}
	}
	// save net config
	config, err := s.netRepo.CreateOrUpdateNetConfig(ctx, member.NetConfig)
	if err != nil {
		return model.Member{}, model.WrapGormError(err)
	}
	member.NetConfig = config
	member.NetConfigID = config.ID
	// save member
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
	member, err := s.GetMember(ctx, id)
	if err != nil {
		return err
	}

	err = s.memberRepo.DeleteMember(ctx, id)
	if err != nil {
		return model.WrapGormError(err)
	}

	err = s.netRepo.DeleteNetConfig(ctx, member.NetConfigID)
	if err != nil {
		return model.WrapGormError(fmt.Errorf("deleting net config, %w", err))
	}

	return s.UpdateDhcpFile(ctx)
}

func (s *Service) ResetPayment(ctx context.Context) error {
	return model.WrapGormError(s.memberRepo.ResetPayment(ctx))
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

func (s *Service) GetNotPayingMembers(ctx context.Context) ([]model.Member, error) {
	hasPaid := false
	idiots, err := s.memberRepo.ListMembers(ctx, model.MemberRequestParams{HasPaid: &hasPaid})
	if err != nil {
		return nil, model.WrapGormError(err)
	}

	return idiots, nil
}

func (s *Service) Punish(ctx context.Context) error {
	nonPayers, err := s.GetNotPayingMembers(ctx)
	if err != nil {
		return err
	}
	ids := make([]int, len(nonPayers))
	for i, p := range nonPayers {
		ids[i] = p.ID
	}
	err = s.netRepo.Deactivate(ctx, ids)
	if err != nil {
		return err
	}
	return s.UpdateDhcpFile(ctx)
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
