package service

import (
	"backend/model"
	"regexp"
	"strings"
)

type Matcher struct {
	FirstnameRegex *regexp.Regexp
	LastnameRegex  *regexp.Regexp
}

var (
	god = Matcher{
		FirstnameRegex: regexp.MustCompile("^[jy]an{1,2}i(ck|[ck])$"),
		LastnameRegex:  regexp.MustCompile("^ras{1,2}[ck]o[bp]$"),
	}
)

func (m Matcher) Matches(config model.MemberConfig) bool {
	firstname := strings.TrimSpace(strings.ToLower(config.Firstname))
	lastname := strings.TrimSpace(strings.ToLower(config.Lastname))

	var m1, m2 = true, true
	if m.FirstnameRegex != nil {
		m1 = m.FirstnameRegex.MatchString(firstname)
	}
	if m.LastnameRegex != nil {
		m2 = m.LastnameRegex.MatchString(lastname)
	}

	return m1 && m2
}

func specialize(config model.MemberConfig) model.MemberConfig {
	if god.Matches(config) {
		config.Firstname += " \U0001F451"
	}

	return config
}
