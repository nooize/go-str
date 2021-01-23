package str

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func (s Str) Url() *url.URL{
	if v, err := url.ParseRequestURI(s.value); err == nil {
		return v
	}
	return nil
}

func (s Str) CamelToSnakeCase() Str {
	s.value = matchFirstCap.ReplaceAllString(s.value, "${1}_${2}")
	s.value = matchAllCap.ReplaceAllString(s.value, "${1}_${2}")
	s.value = strings.ToLower(s.value)
	return s
}

func (s Str) ParseToIntSlice(separator string) (result []int) {
	result = make([]int, 0)
	if len(separator) == 0 {
		return
	}
	for _, v := range strings.Split(s.value, separator) {
		if n, err := strconv.Atoi(v); err == nil {
			result = append(result, n)
		}
	}
	return
}
