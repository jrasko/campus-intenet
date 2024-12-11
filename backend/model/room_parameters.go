package model

import (
	"strings"

	"gorm.io/gorm"
)

type RoomRequestParams struct {
	IsOccupied *bool
	Blocks     []string
	Search     string
	HasPaid    *bool
	Disabled   *bool
	WG         *string
}

var (
	searchableFields = []string{
		"wg",
		"number",
		`"Member".firstname`,
		`"Member".lastname`,
		`"Member".email`,
		`"Member".phone`,
		`"Member".comment`,
		`mac`,
		`ip`,
	}
)

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

	if r.Search != "" {
		db = db.Where(buildSearchQuery(), map[string]any{"s": "%" + r.Search + "%"})
	}
	if r.WG != nil {
		db = db.Where(`wg = ?`, r.WG)
	}
	if r.HasPaid != nil {
		db = db.Where(`"Member".has_paid = ?`, *r.HasPaid)
	}
	if r.Disabled != nil {
		db = db.Where(`disabled = ?`, *r.Disabled)
	}

	return db.Order("number")
}

func buildSearchQuery() string {
	var builder strings.Builder
	for i, field := range searchableFields {
		builder.WriteString(field + " ILIKE @s")
		if i < len(searchableFields)-1 {
			builder.WriteString(" OR ")
		}
	}
	return builder.String()
}
