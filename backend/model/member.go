package model

import (
	"net"
	"strings"
	"time"
)

type Member struct {
	ID         int    `json:"id" gorm:"primaryKey,autoIncrement"`
	Firstname  string `json:"firstname" validate:"required"`
	Lastname   string `json:"lastname" validate:"required"`
	HasPaid    bool   `json:"hasPaid" gorm:"not null"`
	Email      string `json:"email" validate:"omitempty,email"`
	Phone      string `json:"phone"`
	Comment    string `json:"comment"`
	LastEditor string `json:"lastEditor"`

	RoomNr string `json:"-" gorm:"unique;not null"`
	Room   Room   `json:"room" gorm:"foreignKey:RoomNr;references:Number"`

	DhcpID     int        `json:"-" gorm:"unique"`
	DhcpConfig DhcpConfig `json:"dhcpConfig" gorm:"foreignKey:DhcpID"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ReducedMember struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (c *Member) ToReduced() ReducedMember {
	return ReducedMember{
		Firstname: c.Firstname,
		Lastname:  c.Lastname,
	}
}

func (c *Member) Sanitize() {
	mac, _ := net.ParseMAC(c.DhcpConfig.Mac)
	c.DhcpConfig.Mac = strings.ToUpper(mac.String())
}
