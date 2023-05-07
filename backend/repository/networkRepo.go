package repository

import (
	"backend/model"
	"context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(dsn string) (NetworkRepository, error) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return NetworkRepository{}, err
	}
	return NetworkRepository{db: db}, nil
}

type NetworkRepository struct {
	db *gorm.DB
}

func (nr NetworkRepository) UpdateNetworkConfig(ctx context.Context, conf model.NetworkConfig) (model.NetworkConfig, error) {
	err := nr.db.
		WithContext(ctx).
		Save(&conf).
		Error
	if err != nil {
		return model.NetworkConfig{}, err
	}
	return conf, nil
}

func (nr NetworkRepository) GetAllNetworkConfigs(ctx context.Context) ([]model.NetworkConfig, error) {
	var configs []model.NetworkConfig
	err := nr.db.
		WithContext(ctx).
		Order("lastname, firstname").
		Find(&configs).
		Error
	if err != nil {
		return []model.NetworkConfig{}, err
	}
	return configs, nil
}

func (nr NetworkRepository) GetNetworkConfig(ctx context.Context, mac string) (model.NetworkConfig, error) {
	config := model.NetworkConfig{}
	err := nr.db.
		WithContext(ctx).
		Where("mac = ?", mac).
		Find(&config).
		Error
	if err != nil {
		return model.NetworkConfig{}, err
	}
	return config, nil
}

func (nr NetworkRepository) DeleteNetworkConfig(ctx context.Context, mac string) error {
	c := model.NetworkConfig{Mac: mac}
	return nr.db.
		WithContext(ctx).
		Delete(&c).
		Error
}

func (nr NetworkRepository) ResetPayment(ctx context.Context) error {
	return nr.db.
		WithContext(ctx).
		Table("network_configs").
		Where("true").
		Updates(map[string]interface{}{"has_paid": false}).
		Error
}
