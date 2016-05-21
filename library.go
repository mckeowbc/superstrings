package superstrings

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

type Stringer struct {
	Pattern *regexp.Regexp
	MinLen  uint
}

type FoundString struct {
	Str    string
	Offset uint64
}

func NewStringer(charsets []string, min_len uint) *Stringer {
	s := new(Stringer)
	s.Pattern = regexp.MustCompile(fmt.Sprintf(`[[:print:][:blank:]\p{Arabic}\x{200E}\x{200F}]{%d,}`, min_len))
	s.MinLen = min_len
	return s
}

func (v *Stringer) GetStrings(buffer []byte, offset uint64) []FoundString {
	max_cap := v.MinLen / uint(len(buffer))
	out := make([]FoundString, 0, max_cap)

	match := v.Pattern.FindAllIndex(buffer, -1)

	for _, m := range match {
		s := buffer[m[0]:m[1]]

		if utf8.Valid(s) {
			out = append(out, FoundString{string(s), uint64(m[0]) + offset})
		}
	}

	return out
}

func (v FoundString) String() string {
	return fmt.Sprintf("%d\t%s", v.Offset, v.Str)
}
