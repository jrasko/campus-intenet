package model

import (
	"strings"

	"gorm.io/gorm"
)

type MemberRequestParams struct {
	Search   string
	HasPaid  *bool
	Disabled *bool
	WG       *string
}

var (
	searchableFields = []string{
		"firstname",
		"lastname",
		"email",
		"phone",
		"comment",
		"room_nr",
		`"NetConfig".mac`,
		`"NetConfig".ip`,
		`"Room".wg`,
	}
)

func (r MemberRequestParams) Apply(db *gorm.DB) *gorm.DB {
	if r.Search != "" {
		db = db.Where(buildSearchQuery(), map[string]any{"s": "%" + r.Search + "%"})
	}
	if r.WG != nil {
		db = db.Where(`"Room".wg = ?`, r.WG)
	}
	if r.HasPaid != nil {
		db = db.Where("has_paid = ?", *r.HasPaid)
	}
	if r.Disabled != nil {
		db = db.Where("disabled = ?", *r.Disabled)
	}

	return db.Order("room_nr, lastname, firstname")
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
