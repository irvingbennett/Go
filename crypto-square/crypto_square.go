package cryptosquare

import (
	"fmt"
	"math"
	"strings"
)

// const testVersion = 2

func cleanUp(r rune) rune {
	if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
		return r
	}
	return -1
}

// Encode returns an encoded string
func Encode(pt string) (et string) {
	pt = strings.ToLower(pt)
	pt = strings.Map(cleanUp, pt)
	numCols := int(math.Ceil(math.Sqrt(float64(len(pt)))))
	cols := make([]string, numCols)
	for i, r := range pt {
		fmt.Sprintf("'%s' ", string(r))
		cols[i%numCols] += strings.TrimSpace(string(r))
	}
	s := strings.Join(cols, " ")
	r := int(len(pt) / numCols)
	if r*numCols < len(pt) {
		r++
	}
	fmt.Println(pt, "Len:", len(pt), "Rows:", r, "Cols:", numCols)
	if len(pt) < numCols*r {
		return s + " "
	} else {
		return s
	}

}
