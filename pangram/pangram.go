// Package pangram does...
package pangram

import (
	"strings"
)

// IsPangram returns true if a sentence is a pangram (contains every letter of the alphabet)
func IsPangram(sentence string) (is bool) {
	s := strings.ToLower(sentence)
	// fmt.Println(s)
	// is = false
	letters := make(map[byte]bool)
	for x := 0; x <= int(byte('z')-byte('a')); x++ {
		letters[byte('a'+x)] = false
	}
	for _, l := range s {
		// fmt.Printf("%v", l)
		letters[byte(l)] = true
	}
	for _, b := range letters {
		if !b {
			is = false
			return
		}
	}
	is = true

	// fmt.Println(letters)
	return
}
