package model

import (
	"time"
)

const IPValidation = "required,ipv4"

type NetConfig struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Mac          string `json:"mac" validate:"required,mac" gorm:"unique;not null"`
	IP           string `json:"ip" validate:"omitempty,ipv4" gorm:"unique;not null"`
	Manufacturer string `json:"manufacturer" validate:"omitempty,len=0"`
	Disabled     bool   `json:"disabled" gorm:"not null"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
