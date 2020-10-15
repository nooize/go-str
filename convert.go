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
	n := Str{ value: matchFirstCap.ReplaceAllString(s.value, "${1}_${2}") }
	n.value = matchAllCap.ReplaceAllString(n.value, "${1}_${2}")
	n.value = strings.ToLower(n.value)
	return n
}
