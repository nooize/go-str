package str

import (
	"strings"
	"unicode"
)

func (s Str) HasPrefix(prefix string) bool {
	return strings.HasPrefix(s.value, prefix)
}

func (s Str) HasSuffix(suffix string) bool {
	return strings.HasSuffix(s.value, suffix)
}

func (s Str) Replace(old, new string, n int) Str {
	return Str{&}
}

func (s Str) ReplaceAll(old, new string) Str {
	return Str{strings.ReplaceAll(s.value, old, new)}
}

func (s Str) Title() Str {
	return Str{strings.Title(s.value)}
}

func (s Str) ToLower() Str {
	return Str{strings.ToLower(s.value)}
}

func (s Str) ToLowerSpecial(c unicode.SpecialCase) Str {
	return Str{strings.ToLowerSpecial(c, s.value)}
}

func (s Str) ToUpper() Str {
	return Str{strings.ToUpper(s.value)}
}

func (s Str) ToUpperSpecial(c unicode.SpecialCase) Str {
	return Str{strings.ToUpperSpecial(c, s.value)}
}

func (s Str) TrimSpace() Str {
	return Str{strings.TrimSpace(s.value)}
}
