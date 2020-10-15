package str

import (
	"net/url"
)

func (s Str) IsUrl() bool {
	url, err := url.ParseRequestURI(s.value)
	if err != nil {
		return false
	}
	return url != nil && url.IsAbs()
}
