package model

import (
	"strings"

	"gorm.io/gorm"
)

type RequestParams struct {
	Search string
}

var (
	searchableFields = []string{
		"firstname",
		"lastname",
		"mac",
		"room_nr",
		"wg",
		"email",
		"phone",
		"ip",
	}
)

func (r RequestParams) Apply(db *gorm.DB) *gorm.DB {
	if r.Search != "" {
		db = db.Where(buildSearchQuery(), map[string]any{"s": "%" + r.Search + "%"})
	}

	db.Order("lastname, firstname")
	return db
}

func buildSearchQuery() string {
	var builder strings.Builder
	for i, field := range searchableFields {
		builder.WriteString("lower(")
		builder.WriteString(field)
		builder.WriteString(") LIKE lower(@s)")
		if i < len(searchableFields)-1 {
			builder.WriteString(" OR ")

		}
	}
	return builder.String()
}
