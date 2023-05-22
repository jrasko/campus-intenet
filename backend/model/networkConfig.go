package model

type MemberConfig struct {
	ID        int    `json:"id" gorm:"primaryKey,autoIncrement"`
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Mac       string `json:"mac" validate:"required,mac" gorm:"unique"`
	RoomNr    string `json:"roomNr" validate:"required" gorm:"unique"`
	HasPaid   bool   `json:"hasPaid"`
	WG        string `json:"wg"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	IP        string `json:"ip" gorm:"unique"`
}
