package repository

import (
	"backend/model"
	"context"

	"gorm.io/gorm"
)

const memberTable = "members"

type MemberRepository struct {
	db *gorm.DB
}

func (mr MemberRepository) CreateOrUpdateMember(ctx context.Context, conf model.Member) (model.Member, error) {
	err := mr.db.
		WithContext(ctx).
		Save(&conf).
		Error
	return conf, err
}

func (mr MemberRepository) ListMembers(ctx context.Context, params model.MemberRequestParams) ([]model.Member, error) {
	var configs []model.Member
	db := mr.db.
		WithContext(ctx).
		InnerJoins("Room").
		InnerJoins("NetConfig")

	db = params.Apply(db)
	err := db.
		Find(&configs).
		Error
	return configs, err
}

func (mr MemberRepository) GetMember(ctx context.Context, id int) (model.Member, error) {
	config := model.Member{}
	err := mr.db.
		WithContext(ctx).
		InnerJoins("Room").
		InnerJoins("NetConfig").
		First(&config, id).
		Error
	return config, err
}

func (mr MemberRepository) DeleteMembers(ctx context.Context, id int) error {
	return mr.db.
		WithContext(ctx).
		Delete(&model.Member{}, id).
		Error
}

func (mr MemberRepository) ResetPayment(ctx context.Context) error {
	return mr.db.
		WithContext(ctx).
		Table(memberTable).
		Where("true").
		Updates(map[string]interface{}{"has_paid": false}).
		Error
}
