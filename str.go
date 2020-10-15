package str

type Str struct {
	value string
}

func New(input string) Str {
	return Str{ value: input }
}

func (s Str) String() string {
	return s.value
}

func (s Str) S() string {
	return s.value
}


