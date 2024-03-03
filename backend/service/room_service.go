package service

import (
	"backend/model"
	"context"
)

func (s *Service) ListRooms(ctx context.Context, params model.RoomRequestParams) ([]model.Room, error) {
	rooms, err := s.roomRepo.ListRooms(ctx, params)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
