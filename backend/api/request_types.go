package api

import (
	"backend/model"
)

type requestMember struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	HasPaid   bool   `json:"hasPaid" gorm:"not null"`
	Email     string `json:"email" validate:"omitempty,email"`
	Phone     string `json:"phone"`
	Comment   string `json:"comment"`
	MovedIn   string `json:"movedIn"`

	RoomNr     string          `json:"roomNr" gorm:"unique;not null"`
	DhcpConfig model.NetConfig `json:"dhcpConfig" gorm:"foreignKey:NetConfigID"`
}

func (r requestMember) toModel() model.Member {
	return model.Member{
		Firstname: r.Firstname,
		Lastname:  r.Lastname,
		HasPaid:   r.HasPaid,
		Email:     r.Email,
		Phone:     r.Phone,
		Comment:   r.Comment,
		RoomNr:    r.RoomNr,
		NetConfig: r.DhcpConfig,
		MovedIn:   r.MovedIn,
	}
}
