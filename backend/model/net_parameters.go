package model

import (
	"gorm.io/gorm"
)

type NetworkRequestParams struct {
	Servers  *bool
	Disabled *bool
}

func (r NetworkRequestParams) Apply(db *gorm.DB) *gorm.DB {
	if r.Disabled != nil {
		db = db.Where("disabled = ?", r.Disabled)
	}
	if r.Servers != nil {
		if *r.Servers {
			db = db.Where(`"Member".id IS NULL`)
		} else {
			db = db.Where(`"Member".id IS NOT NULL`)
		}
	}

	return db.Order("ip")
}
