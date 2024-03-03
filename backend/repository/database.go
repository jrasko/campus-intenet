package repository

import (
	"backend/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(dsn string) (MemberRepository, error) {
	log.Println("[INFO] Connecting to Database...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return MemberRepository{}, err
	}
	log.Println("[INFO] Successfully connected to DB")

	// check if db has needed table
	if err = db.AutoMigrate(&model.Member{}, &model.Room{}); err != nil {
		return MemberRepository{}, fmt.Errorf("could not create table: %w", err)
	}
	return MemberRepository{db: db}, nil
}
