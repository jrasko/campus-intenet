package repository

import (
	"backend/model"
	"context"
)

func (mr MemberRepository) GetEnabledNets(ctx context.Context) ([]model.DhcpConfig, error) {
	var members []model.DhcpConfig
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
