package model

type NetworkConfig struct {
	Name    string `json:"name" validate:"required"`
	Mac     string `json:"mac" validate:"required,mac" gorm:"primaryKey"`
	RoomNr  string `json:"roomNr"`
	HasPaid bool   `json:"hasPaid"`
	WG      string `json:"wg"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}
