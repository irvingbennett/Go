// Package isbn verifies a code
package isbn

import (
	"unicode"
)

// IsValidISBN (test.isbn) checks the ISBN
func IsValidISBN(isbn string) (valid bool) {
	// fmt.Println(isbn)
	n := 0
	x := 10
	last := ' '
	for _, d := range isbn {
		if unicode.IsDigit(d) || unicode.ToUpper(d) == 'X' {
			if unicode.IsDigit(d) {
				n += (int(d) - 48) * x
				// fmt.Printf("%d * %d = %d, n = %d\n", (int(d) - 48), x, (int(d)-47)*x, n)
				x--
			} else {
				n += 10
				x--
				// fmt.Println("Is X, n is", n)
			}
		} else if unicode.IsLetter(d) && x > 0 {
			valid = false
			return
		}
		last = d
	}

	if !unicode.IsDigit(last) && !(unicode.ToUpper(last) == 'X') || x != 0 {
		valid = false
		return
	}
	if n%11 == 0 {
		valid = true
	}
	// fmt.Printf("Last: %s, n: %d, x: %d, valid: %v\n", string(last), n, x, valid)
	return
}
