package repository

import (
	"backend/model"
	"context"
)

const netTable = "net_configs"

func (mr MemberRepository) ListRooms(ctx context.Context, params model.RoomRequestParams) ([]model.Room, error) {
	var rooms []model.Room

	tx := mr.db.
		WithContext(ctx).
		Joins("Member")
	tx = params.Apply(tx)
	err := tx.
		Order("number").
		Find(&rooms).
		Error
	return rooms, err
}

func (mr MemberRepository) GetRoom(ctx context.Context, number string) (model.Room, error) {
	var room model.Room
	err := mr.db.
		WithContext(ctx).
		First(&room, number).
		Error
	return room, err
}
