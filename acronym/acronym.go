package acronym

import (
	"regexp"
	"strings"
)

func removeSpace(s string) string {
	space := regexp.MustCompile(`[\s+\-_]`)
	str := space.ReplaceAllString(s, " ")
	return str
}

// Abbreviate should have a comment documenting it.
func Abbreviate(str string) string {
	s := removeSpace(str)
	initials := make([]byte, 20)
	words := strings.Split(s, " ")
	i := 0
	for w := range words {
		upper := strings.ToUpper(words[w])
		if len(upper) > 0 {
			initials[i] = upper[0]
			i++
		}
	}
	return string(initials[:i])
}
