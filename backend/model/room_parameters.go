package model

import (
	"gorm.io/gorm"
)

type RoomRequestParams struct {
	IsOccupied *bool
	Blocks     []string
}

func (r RoomRequestParams) Apply(db *gorm.DB) *gorm.DB {
	if len(r.Blocks) > 0 {
		db = db.Where("block in ?", r.Blocks)
	}
	if r.IsOccupied != nil {
		if *r.IsOccupied {
			db = db.Where(`"Member".id IS NOT NULL`)
		} else {
			db = db.Where(`"Member".id IS NULL`)
		}
	}

	return db.Order("number")
}
