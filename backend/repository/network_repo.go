package repository

import (
	"backend/model"
	"context"
	"fmt"
	"log"

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
	log.Println("Connecting to Database...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return MemberRepository{}, err
	}
	log.Println("Successfully connected to DB")

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

func (mr MemberRepository) UpdateMemberConfig(ctx context.Context, conf model.MemberConfig) (model.MemberConfig, error) {
	err := mr.db.
		WithContext(ctx).
		Save(&conf).
		Error
	return conf, err
}

func (mr MemberRepository) GetAllMemberConfigs(ctx context.Context, params model.RequestParams) ([]model.MemberConfig, error) {
	var configs []model.MemberConfig
	db := mr.db.WithContext(ctx)
	db = params.Apply(db)
	err := db.
		Find(&configs).
		Error
	return configs, err
}

func (mr MemberRepository) GetEnabledUsers(ctx context.Context) ([]model.MemberConfig, error) {
	var members []model.MemberConfig
	err := mr.db.
		WithContext(ctx).
		Where("disabled = false").
		Find(&members).
		Error
	return members, err
}

func (mr MemberRepository) GetMemberConfig(ctx context.Context, id int) (model.MemberConfig, error) {
	config := model.MemberConfig{}
	err := mr.db.
		WithContext(ctx).
		Where("id = ?", id).
		First(&config).
		Error
	return config, err
}

func (mr MemberRepository) DeleteMemberConfig(ctx context.Context, id int) error {
	return mr.db.
		WithContext(ctx).
		Delete(&model.MemberConfig{}, id).
		Error
}

func (mr MemberRepository) ResetPayment(ctx context.Context) error {
	return mr.db.
		WithContext(ctx).
		Table(memberTable).
		Where("true").
		Updates(map[string]interface{}{"has_paid": false}).
		Error
}

func (mr MemberRepository) GetAllIPs(ctx context.Context) ([]string, error) {
	var ips []string
	err := mr.db.
		WithContext(ctx).
		Table(memberTable).
		Select("ip").
		Order("ip").
		Find(&ips).
		Error
	return ips, err
}

func (mr MemberRepository) GetNonPayingMembers(ctx context.Context) ([]model.MemberConfig, error) {
	var members []model.MemberConfig
	err := mr.db.
		WithContext(ctx).
		Select("firstname", "lastname").
		Where("has_paid = false").
		Find(&members).
		Error
	return members, err
}
