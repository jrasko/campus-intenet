package model

import (
	"time"
)

type NetConfig struct {
	ID           int    `gorm:"primaryKey"`
	Mac          string `json:"mac" validate:"required,mac" gorm:"unique;not null"`
	IP           string `json:"ip" validate:"omitempty,ipv4" gorm:"unique;not null"`
	Manufacturer string `json:"manufacturer" validate:"omitempty,len=0"`
	Disabled     bool   `json:"disabled" gorm:"not null"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
