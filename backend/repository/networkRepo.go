package repository

import (
	"backend/model"

	"golang.org/x/net/context"
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

func (nr NetworkRepository) UpdateNetworkConfig(ctx context.Context, conf model.NetworkConfig) error {
	return nr.db.
		WithContext(ctx).
		Save(&conf).
		Error
}
