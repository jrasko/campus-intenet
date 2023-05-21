package repository

import (
	"backend/model"
	"context"
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("db entry not found")

const (
	memberTable = "member_configs"
)

type AllocatedIP struct {
	IP string
}

func New(dsn string) (NetworkRepository, error) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return NetworkRepository{}, err
	}
	if !db.Migrator().HasTable(memberTable) {
		return NetworkRepository{}, errors.New("missing table")
	}
	return NetworkRepository{db: db}, nil
}

type NetworkRepository struct {
	db *gorm.DB
}

func (nr NetworkRepository) UpdateNetworkConfig(ctx context.Context, conf model.MemberConfig) (model.MemberConfig, error) {
	err := nr.db.
		WithContext(ctx).
		Save(&conf).
		Error
	return conf, err
}

func (nr NetworkRepository) GetAllNetworkConfigs(ctx context.Context) ([]model.MemberConfig, error) {
	var configs []model.MemberConfig
	err := nr.db.
		WithContext(ctx).
		Order("lastname, firstname").
		Find(&configs).
		Error
	return configs, err
}

func (nr NetworkRepository) GetAllMacs(ctx context.Context) ([]string, error) {
	var macs []string
	err := nr.db.
		WithContext(ctx).
		Table(memberTable).
		Select("mac").
		Find(&macs).
		Error
	return macs, err
}

func (nr NetworkRepository) GetNetworkConfig(ctx context.Context, id int) (model.MemberConfig, error) {
	config := model.MemberConfig{}
	err := nr.db.
		WithContext(ctx).
		Where("id = ?", id).
		First(&config).
		Error
	if err == gorm.ErrRecordNotFound {
		return model.MemberConfig{}, ErrNotFound
	} else if err != nil {
		return model.MemberConfig{}, err
	}
	return config, nil
}

func (nr NetworkRepository) DeleteNetworkConfig(ctx context.Context, id int) error {
	return nr.db.
		WithContext(ctx).
		Delete(&model.MemberConfig{}, id).
		Error
}

func (nr NetworkRepository) ResetPayment(ctx context.Context) error {
	return nr.db.
		WithContext(ctx).
		Table(memberTable).
		Where("true").
		Updates(map[string]interface{}{"has_paid": false}).
		Error
}

func (nr NetworkRepository) GetAllIPs(ctx context.Context) ([]string, error) {
	var ips []string
	err := nr.db.
		WithContext(ctx).
		Table(memberTable).
		Select("ip").
		Order("ip").
		Find(&ips).
		Error
	return ips, err
}
