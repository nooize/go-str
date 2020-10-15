package str

import (
	"regexp"
	"strings"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func (s Str) CamelToSnakeCase() Str {
	s.value = matchFirstCap.ReplaceAllString(s.value, "${1}_${2}")
	s.value = matchAllCap.ReplaceAllString(s.value, "${1}_${2}")
	s.value = strings.ToLower(s.value)
	return s
}
