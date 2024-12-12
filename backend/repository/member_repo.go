package repository

import (
	"backend/model"
	"context"

	"gorm.io/gorm"
)

const memberTable = "members"

type Repository struct {
	db *gorm.DB
}

func (r Repository) CreateOrUpdateMember(ctx context.Context, conf model.Member) (model.Member, error) {
	err := r.db.
		WithContext(ctx).
		Omit("Room").
		Save(&conf).
		Error
	return conf, err
}

func (r Repository) GetMember(ctx context.Context, id int) (model.Member, error) {
	config := model.Member{}
	err := r.db.
		WithContext(ctx).
		InnerJoins("Room").
		InnerJoins("NetConfig").
		First(&config, id).
		Error
	return config, err
}

func (r Repository) DeleteMember(ctx context.Context, id int) error {
	return r.db.
		WithContext(ctx).
		Delete(&model.Member{}, id).
		Error
}

func (r Repository) ResetPayment(ctx context.Context) error {
	return r.db.
		WithContext(ctx).
		Table(memberTable).
		Where("true").
		Updates(map[string]any{"has_paid": false}).
		Error
}
