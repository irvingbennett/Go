// Package diamond prints a diamond of letters
package diamond

import (
	"bytes"
	"fmt"
	"strings"
)

// Gen generates a diamond from a letter
func Gen(c byte) (s string, err error) {
	// fmt.Println(string(c))
	n := bytes.ToUpper([]byte{c})
	if c < 'A' || c > 'Z' {
		err = fmt.Errorf("input out of range")
		return
	}
	l := n[0] - byte('A')
	diamond := make([][]byte, int(l)*2+2)
	b := make([]string, 0)
	p1 := int(l)
	p2 := p1
	// fmt.Println(l)
	a := make([]byte, int(l)+1)
	for x := 0; x <= int(l); x++ {
		a[x] = byte(65 + x)
	}
	// fmt.Println(string(a))
	letter := 0
	out := true
	for x := 0; x <= int(l)*2; x++ {
		diamond[x] = make([]byte, int(l)*2+2)
		for y := 0; y <= int(l)*2; y++ {
			diamond[x][y] = ' '
			if y == p1 {
				diamond[x][y] = a[letter]
			}
			if y == p2 {
				diamond[x][y] = a[letter]
			}
		}
		// fmt.Println(string(diamond[x]))
		if out {
			letter++
			p1--
			p2++
		} else {
			letter--
			p1++
			p2--
		}

		if p1 == -1 {
			out = false
			p1 = 1
			p2 -= 2
			letter -= 2
		}
	}
	for x := range diamond {
		if len(diamond[x]) > 0 {
			b = append(b, string(diamond[x]))
		}
	}
	s = strings.Join(b, "\n")
	s = s + "\n"
	// fmt.Println(s)
	return
}
