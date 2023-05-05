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
	configs := []model.NetworkConfig{}
	err := nr.db.
		WithContext(ctx).
		Find(&configs).
		Order("lastname").
		Error
	if err != nil {
		return []model.NetworkConfig{}, err
	}
	return configs, nil
}
