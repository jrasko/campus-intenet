package model

type Room struct {
	Number string  `json:"roomNr" gorm:"primaryKey"`
	WG     string  `json:"wg"`
	Block  string  `json:"block"`
	Member *Member `json:"member" gorm:"foreignKey:RoomNr;references:Number"`
}
