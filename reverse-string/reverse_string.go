// Package reverse reverses a string and returns it
package reverse

import (
	"unicode/utf8"
)

// Reverse reverses a string and returns it
func Reverse(s string) string {
	b := []byte(s)
	runes := make([]rune, len(b))
	x := 0
	// fmt.Printf("%s, Len: %d -- Runes: %d \n", s, len(s), utf8.RuneCountInString(s))
	for len(b) > 0 {
		r, size := utf8.DecodeLastRune(b)
		runes[x] = r
		x++
		b = b[:len(b)-size]
	}

	// fmt.Printf("%s -> %s | runes: %d\n\n", s, string(runes), len(runes))
	return string(runes[:x])
}
