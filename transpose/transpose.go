// Package transpose transposes a matrix
package transpose

import (
	"bytes"
)

// Transpose transposes a matrix
func Transpose(m []string) []string {
	if len(m) == 0 {
		return []string{}
	}
	r := make([][]byte, testMatrix(m))
	for x := range r {
		r[x] = make([]byte, len(m))
	}
	for y, s := range m {
		for x, e := range s {
			r[x][y] = byte(e)
		}
	}
	t := []string{}
	for i := 0; i < len(r); i++ {
		row := bytes.TrimRight(r[i], "\x00")
		row = bytes.ReplaceAll(row, []byte("\x00"), []byte(" "))
		t = append(t, string(row))
	}
	return t
}

func testMatrix(in []string) int {
	l := len(in[0])
	for i := 0; i < len(in); i++ {
		if l < len(in[i]) {
			l = len(in[i])
		}
	}
	return l
}
