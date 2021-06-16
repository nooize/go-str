package str

import (
	"strings"
	"unicode"
)


func (s Str) Add(v Str) Str {
	s.value += v.S()
	return s
}

func (s Str) AddString(v string) Str {
	s.value += v
	return s
}

func (s Str) HasPrefix(prefix string) bool {
	return strings.HasPrefix(s.value, prefix)
}

func (s Str) HasSuffix(suffix string) bool {
	return strings.HasSuffix(s.value, suffix)
}

func (s Str) Replace(old, new string, n int) Str {
	s.value = strings.Replace(s.value, old, new, n)
	return s
}

func (s Str) ReplaceAll(old, new string) Str {
	s.value = strings.ReplaceAll(s.value, old, new)
	return s
}

func (s Str) Title() Str {
	s.value = strings.Title(s.value)
	return s
}

func (s Str) ToLower() Str {
	s.value = strings.ToLower(s.value)
	return s
}

func (s Str) ToLowerSpecial(c unicode.SpecialCase) Str {
	s.value = strings.ToLowerSpecial(c, s.value)
	return s
}

func (s Str) ToUpper() Str {
	s.value = strings.ToUpper(s.value)
	return s
}

func (s Str) ToUpperSpecial(c unicode.SpecialCase) Str {
	s.value = strings.ToUpperSpecial(c, s.value)
	return s
}

func (s Str) Trim(curset string) Str {
	s.value = strings.Trim(s.value, curset)
	return s
}

func (s Str) TrimSpace() Str {
	s.value = strings.TrimSpace(s.value)
	return s
}
