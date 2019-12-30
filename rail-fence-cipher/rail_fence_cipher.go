// Package railfence does...
package railfence

import (
	"strings"
)

// Encode returns a cypher
func Encode(s string, rails int) (cypher string) {
	b := []byte(strings.ReplaceAll(s, " ", ""))
	r := make([][]byte, rails)
	// fmt.Println(s, string(b), rails)

	row := 0
	down := true
	for x := 0; x < rails; x++ {
		r[x] = make([]byte, 0)
	}
	// fmt.Println(r)
	for x := 0; x < len(b); x++ {
		for y := 0; y < rails; y++ {
			if y == row {
				r[y] = append(r[y], b[x])
			} else {
				r[y] = append(r[y], '.')
			}
			// fmt.Println(r)
		}

		if down == true {
			row++
		} else {
			row--
		}
		if row == rails {
			down = false
			row = rails - 2
		}
		if row == -1 {
			down = true
			row += 2
		}

	}
	code := make([]byte, 0)
	for x := 0; x < rails; x++ {
		for y := 0; y < len(r[x]); y++ {
			l := r[x][y]
			if l != '.' {
				code = append(code, l)
			}

		}
	}
	cypher = string(code)
	// fmt.Println(cypher)
	return
}

// Decode returns a cypher
func Decode(s string, rails int) (cypher string) {
	b := []byte(strings.ReplaceAll(s, " ", ""))
	r := make([][]byte, rails)
	// fmt.Println(s, string(b), rails)

	row := 0
	down := true
	for x := 0; x < rails; x++ {
		r[x] = make([]byte, 0)
	}
	// fmt.Println(r)
	for x := 0; x < len(b); x++ {
		for y := 0; y < rails; y++ {
			if y == row {
				r[y] = append(r[y], '_')
			} else {
				r[y] = append(r[y], '.')
			}
			// fmt.Println(r)
		}

		if down == true {
			row++
		} else {
			row--
		}
		if row == rails {
			down = false
			row = rails - 2
		}
		if row == -1 {
			down = true
			row += 2
		}

	}
	code := make([]byte, 0)
	idx := 0
	for x := 0; x < rails; x++ {
		for y := 0; y < len(r[x]); y++ {
			l := r[x][y]
			if l == '_' {
				r[x][y] = b[idx]
				idx++
			}

		}
	}

	idx = 0
	for x := 0; x < len(b); x++ {

		for y := 0; y < rails; y++ {
			l := r[y][idx]
			if l != '.' {
				code = append(code, l)
			}
		}
		idx++
	}
	cypher = string(code)
	// fmt.Println(cypher)
	return
}
