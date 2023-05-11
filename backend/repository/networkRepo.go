package repository

import (
	"backend/model"
	"context"
	"errors"
	"net"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("db entry not found")

func New(dsn string) (NetworkRepository, error) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return NetworkRepository{}, err
	}
	if !db.Migrator().HasTable("network_configs") {
		return NetworkRepository{}, errors.New("missing table")
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
	if err == gorm.ErrRecordNotFound {
		return model.NetworkConfig{}, ErrNotFound
	} else if err != nil {
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

const (
	ip_table = "allocated_ips"
)

func (nr NetworkRepository) GetAllIPs(ctx context.Context) ([]net.IP, error) {
	var ips []net.IP
	err := nr.db.
		WithContext(ctx).
		Table(ip_table).
		Order("ip").
		Find(ips).
		Error
	return ips, err
}

func (nr NetworkRepository) AddIP(ctx context.Context, ip net.IP) error {
	return nr.db.
		WithContext(ctx).
		Table(ip_table).
		Save(ip).
		Error
}

func (nr NetworkRepository) RemoveIP(ctx context.Context, ip net.IP) error {
	return nr.db.
		WithContext(ctx).
		Table(ip_table).
		Delete(ip).
		Error
}
