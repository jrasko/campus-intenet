package model

type NetworkConfig struct {
	Name    string `json:"name" validate:"required"`
	Mac     string `json:"mac" validate:"required,mac" gorm:"primaryKey"`
	RoomNr  string
	HasPaid bool
	WG      string
	Email   string
	Phone   string
}
