package repository

import (
	"backend/model"
	"context"
)

func (mr MemberRepository) CreateOrUpdateNetConfig(ctx context.Context, config model.NetConfig) (model.NetConfig, error) {
	err := mr.db.
		WithContext(ctx).
		Save(&config).
		Error
	return config, err
}

func (mr MemberRepository) GetNetConfig(ctx context.Context, id int) (model.NetConfig, error) {
	var member model.NetConfig
	err := mr.db.
		WithContext(ctx).
		First(&member, id).
		Error
	return member, err
}

func (mr MemberRepository) ListNetConfigs(ctx context.Context, params model.NetworkRequestParams) ([]model.NetConfig, error) {
	var members []model.NetConfig
	tx := mr.db.
		WithContext(ctx)

	tx = params.Apply(tx)
	err := tx.
		Find(&members).
		Error
	return members, err
}

func (mr MemberRepository) GetAllIPs(ctx context.Context) ([]string, error) {
	var ips []string
	err := mr.db.
		WithContext(ctx).
		Table(netTable).
		Select("ip").
		Order("ip").
		Find(&ips).
		Error
	return ips, err
}
