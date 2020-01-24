// Package armstrong does...
package armstrong

import (
	"math"
	"strconv"
)

// IsNumber (a int) bool
func IsNumber(a int) (ok bool) {
	// fmt.Println(a)
	n := strconv.Itoa(a)
	power := len(n)
	sum := 0
	for _, c := range n {
		x, good := strconv.Atoi(string(c))
		if good != nil {
			ok = false
			return
		}
		sum += int(math.Pow(float64(x), float64(power)))
	}
	if a == sum {
		ok = true
	}
	return
}
