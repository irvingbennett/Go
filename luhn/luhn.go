package luhn

import (
	"regexp"
	"strconv"
)

func removeSpace(s string) string {
	space := regexp.MustCompile(`\s+`)
	str := space.ReplaceAllString(s, "")
	return str
}

// Valid checks if an input is a valid luhn number
func Valid(s string) bool {
	n := removeSpace(s)
	l := make([]byte, len(n))
	// fmt.Printf("%s --> %s, len: %d", string(l), n, len(l))
	odd := len(n) & 1
	if len(l) <= 1 {
		return false
	}
	p := 0
	result := 0
	for x := 0; x < len(n); x++ {
		number := rune(n[x])
		i, err := strconv.Atoi(string(number))
		if err != nil {
			return false
		}

		if err == nil {
			if (p+odd)%2 == 0 {
				i = i * 2
				if i > 9 {
					i = i - 9
				}
			}
			d := strconv.Itoa(i)
			l[p] = d[0]
			p++
		}
		result += i
	}
	l = l[:p]

	return result%2 == 0

}
