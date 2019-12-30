// Package romannumerals returns the roman numeral of a number
package romannumerals

import (
	"fmt"
)

var (
	m0 = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	m1 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	m2 = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	m3 = []string{"", "M", "MM", "MMM"}
)

// ToRomanNumeral (n int) (roman string, err error)
func ToRomanNumeral(n int) (roman string, err error) {
	if n < 1 || n >= 3001 {
		return "", fmt.Errorf("Bad input")
	}
	// this is efficient in Go.  the seven operands are evaluated,
	// then a single allocation is made of the exact size needed for the result.
	return m3[n%1e4/1e3] +
			m2[n%1e3/1e2] + m1[n%100/10] + m0[n%10],
		nil
}
