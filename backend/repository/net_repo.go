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
	var netConfigs []model.NetConfig
	tx := mr.db.
		WithContext(ctx).
		Joins("LEFT JOIN members ON members.net_config_id = net_configs.id")

	tx = params.Apply(tx)
	err := tx.
		Find(&netConfigs).
		Error
	return netConfigs, err
}
func (mr MemberRepository) DeleteNetConfig(ctx context.Context, id int) error {
	return mr.db.
		WithContext(ctx).
		Delete(&model.NetConfig{}, id).
		Error
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
