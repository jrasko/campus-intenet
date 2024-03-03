package repository

import (
	"backend/model"
	"context"
)

func (mr MemberRepository) SaveNetConfig(ctx context.Context, config model.NetConfig) (model.NetConfig, error) {
	err := mr.db.
		Debug().
		WithContext(ctx).
		Save(&config).
		Error
	return config, err
}

func (mr MemberRepository) GetEnabledNets(ctx context.Context) ([]model.NetConfig, error) {
	var members []model.NetConfig
	err := mr.db.
		WithContext(ctx).
		Where("disabled = false").
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
