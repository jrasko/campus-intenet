package repository

import (
	"backend/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(dsn string) (Repository, error) {
	log.Println("[INFO] Connecting to Database...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return Repository{}, err
	}
	log.Println("[INFO] Successfully connected to DB")

	// check if db has needed table
	if err = db.AutoMigrate(&model.Member{}, &model.Room{}); err != nil {
		return Repository{}, fmt.Errorf("could not create table: %w", err)
	}
	return Repository{db: db}, nil
}
