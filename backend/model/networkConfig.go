package model

type NetworkConf struct {
	Name    string `json:"name" validate:"required"`
	Mac     string `json:"mac" validate:"required,mac"`
	RoomNr  string
	HasPaid bool
	WG      string
	Email   string
	Phone   string
}
