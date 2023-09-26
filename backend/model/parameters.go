package model

import (
	"regexp"
	"slices"
	"strings"

	"gorm.io/gorm"
)

type RequestParams struct {
	Search string
	Order  string `validate:"omitempty,oneof=firstname lastname mac roomNr hasPaid disabled wg email phone ip"`
}

var (
	dbFields = []string{
		"firstname",
		"lastname",
		"mac",
		"room_nr",
		"has_paid",
		"disabled",
		"wg",
		"email",
		"phone",
		"ip",
	}
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

	if r.Order != "" {
		// care! possible sql injection security risk
		column := toGormColumn(r.Order)
		db = db.Order(column)
	}
	db.Order("lastname, firstname")
	return db
}

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
)

func toGormColumn(fieldname string) string {
	column := strings.ToLower(matchFirstCap.ReplaceAllString(fieldname, "${1}_${2}"))
	if slices.Contains(dbFields, column) {
		return column
	}
	return ""
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
