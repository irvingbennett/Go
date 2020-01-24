// Package atbash codes and decodes text
package atbash

import (
	"regexp"
	"strings"
)

// Atbash codes and decodes atbash cyphers
func Atbash(in string) (out string) {
	// fmt.Println(in)
	re := regexp.MustCompile(`[^\w\d]`)
	s := strings.ToLower(re.ReplaceAllString(in, ""))
	o := make([]rune, len(s))

	code := make(map[rune]rune, 0)
	for i, j := int('a'), int('z'); i <= int('z'); i++ {
		code[rune(i)] = rune(j)
		j--
	}
	i := 0
	for _, c := range s {
		o[i] = code[c]
		// fmt.Printf("%s", string(c))
		if c >= '0' && c <= '9' {
			o[i] = c
		}
		i++
	}
	c5 := make([]string, 0)
	for x := 0; x <= len(o); x += 5 {
		y := 5
		if x+5 <= len(o) {
			y = x + 5
		} else {
			y = len(o)
		}

		c5 = append(c5, string(o[x:y]))
	}
	out = strings.TrimSpace(strings.Join(c5, " "))
	// fmt.Println(out)

	return
}
