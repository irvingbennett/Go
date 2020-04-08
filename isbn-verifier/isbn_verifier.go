// Package isbn verifies a code
package isbn

import (
	"fmt"
	"unicode"
)

// IsValidISBN (test.isbn) verifies the code
func IsValidISBN(isbn string) (valid bool) {
	fmt.Println(isbn)
	i := []byte(isbn)
	x := 10
	n := 0
	var last rune
	for _, y := range i {
		if unicode.IsDigit(rune(y)) {
			n = n + (int(y)-47)*x
			x--
		}
		last = rune(y)
	}
	if last == 'x' || last == 'X' {
		n = n + 10
		fmt.Println("Last is X", n, n%11)
	}
	if n%11 == 0 {
		valid = true
	}
	fmt.Printf("X: %d, Valid: %v, Y: %s\n", x, valid, string(last))

	return
}
