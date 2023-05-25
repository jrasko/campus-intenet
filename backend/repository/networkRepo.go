package repository

import (
	"backend/model"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	memberTable = "member_configs"
)

type AllocatedIP struct {
	IP string
}

func New(dsn string) (MemberRepository, error) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return MemberRepository{}, err
	}

	// check if db has needed table
	if !db.Migrator().HasTable(memberTable) {
		if err = db.Migrator().CreateTable(&model.MemberConfig{}); err != nil {
			return MemberRepository{}, fmt.Errorf("could not create table: %w", err)
		}
	}
	return MemberRepository{db: db}, nil
}

type MemberRepository struct {
	db *gorm.DB
}

func (nr MemberRepository) UpdateMemberConfig(ctx context.Context, conf model.MemberConfig) (model.MemberConfig, error) {
	err := nr.db.
		WithContext(ctx).
		Save(&conf).
		Error
	return conf, wrapGormErrors(err)
}

func (nr MemberRepository) GetAllMemberConfigs(ctx context.Context) ([]model.MemberConfig, error) {
	var configs []model.MemberConfig
	err := nr.db.
		WithContext(ctx).
		Order("lastname, firstname").
		Find(&configs).
		Error
	return configs, wrapGormErrors(err)
}

func (nr MemberRepository) GetAllMacs(ctx context.Context) ([]string, error) {
	var macs []string
	err := nr.db.
		WithContext(ctx).
		Table(memberTable).
		Select("mac").
		Find(&macs).
		Error
	return macs, wrapGormErrors(err)
}

func (nr MemberRepository) GetMemberConfig(ctx context.Context, id int) (model.MemberConfig, error) {
	config := model.MemberConfig{}
	err := nr.db.
		WithContext(ctx).
		Where("id = ?", id).
		First(&config).
		Error
	return config, wrapGormErrors(err)
}

func (nr MemberRepository) DeleteMemberConfig(ctx context.Context, id int) error {
	err := nr.db.
		WithContext(ctx).
		Delete(&model.MemberConfig{}, id).
		Error
	return wrapGormErrors(err)
}

func (nr MemberRepository) ResetPayment(ctx context.Context) error {
	err := nr.db.
		WithContext(ctx).
		Table(memberTable).
		Where("true").
		Updates(map[string]interface{}{"has_paid": false}).
		Error
	return wrapGormErrors(err)
}

func (nr MemberRepository) GetAllIPs(ctx context.Context) ([]string, error) {
	var ips []string
	err := nr.db.
		WithContext(ctx).
		Table(memberTable).
		Select("ip").
		Order("ip").
		Find(&ips).
		Error
	return ips, wrapGormErrors(err)
}

func wrapGormErrors(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Error(http.StatusNotFound, err.Error(), "db entry not found")
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		return model.Error(http.StatusConflict, err.Error(), "unique constraint violation")
	}
	return model.Error(http.StatusInternalServerError, err.Error(), "internal server error")
}
