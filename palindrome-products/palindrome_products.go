// Package palindrome does...
package palindrome

import (
	"fmt"
	"strconv"
)

// Product is a list of factors
type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

// Trim removes duplicates and trims slice of products
func trim(p *Product) {
	for x := 0; x < len(p.Factorizations); x++ {
		if p.Factorizations[x] == [2]int{0, 0} {
			p.Factorizations = p.Factorizations[:x]
			// fmt.Println(p)
		}
	}
}

// Products (test.fmin, test.fmax) return the smallest and the largest palindrom within a range with its factors
func Products(min, max int) (pmin, pmax Product, err error) {
	if min > max {
		err = fmt.Errorf("fmin > fmax")
		return
	}
	var newP Product
	fmt.Println(min, max, pmin, pmax)
	p := make(map[int]Product, 32)

	// idx := 0
	for m := min; m <= max; m++ {
		for n := min; n <= max; n++ {
			i := n * m
			if !isPalindrome(i) {
				continue
			}
			// fmt.Println("Is palindrome:", i, "Factors:", n, m)

			newP = Product{Product: i}
			newP.Factorizations = make([][2]int, 32)
			if m <= n {
				newP.Factorizations[0] = [2]int{m, n}
			} else {
				newP.Factorizations[0] = [2]int{n, m}
			}
			if f, ok := p[i]; ok {
				l := len(f.Factorizations)
				for x := 0; x < l; x++ {
					if p[i].Factorizations[x] == [2]int{0, 0} {
						l = x
						break
					}
					// fmt.Println(p[i].Factorizations[x], newP.Factorizations[0])
					if p[i].Factorizations[x] == newP.Factorizations[0] {
						l = x
						break
					}
				}
				// fmt.Println("Second factor:", i, "Len:", l)

				// fmt.Println(p[i], l, newP)
				p[i].Factorizations[l] = newP.Factorizations[0]
			} else {
				p[i] = newP
			}

			if pmin.Product == 0 {
				pmin = newP
			} else if pmin.Product > i {
				pmin = p[i]
			}
			if pmax.Product == 0 {
				pmax = newP
			} else if pmax.Product < i {
				pmax = p[i]
			}

		}

	}
	trim(&pmin)
	trim(&pmax)
	fmt.Println(min, max, pmin, pmax, pmin.Product)
	if pmin.Product == 0 {
		// fmt.Println("No palindromes")
		err = fmt.Errorf("no palindromes")
	}
	return
}
func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func isPalindrome(i int) bool {
	s := strconv.Itoa(i)
	b := []byte(s)
	reverse(b)
	if s == string(b) {
		return true
	}
	return false

}
