package model

type MemberConfig struct {
	ID        int    `json:"id" gorm:"primaryKey,autoIncrement"`
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Mac       string `json:"mac" validate:"required,mac" gorm:"unique"`
	RoomNr    string `json:"roomNr" validate:"required" gorm:"unique"`
	HasPaid   bool   `json:"hasPaid"`
	Disabled  bool   `json:"disabled" gorm:"not null"`
	WG        string `json:"wg"`
	Email     string `json:"email" validate:"omitempty,email"`
	Phone     string `json:"phone"`
	IP        string `json:"ip" gorm:"unique" validate:"omitempty,ipv4"`
}

type ReducedMember struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (c MemberConfig) ToReduced() ReducedMember {
	return ReducedMember{
		Firstname: c.Firstname,
		Lastname:  c.Lastname,
	}
}
