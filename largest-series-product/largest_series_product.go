// Package lsproduct does...
package lsproduct

import (
	"fmt"
	"strconv"
)

// LargestSeriesProduct (test.digits, test.span) returns
func LargestSeriesProduct(series string, span int) (p int, err error) {
	fmt.Println("Series:", series, "Span:", span)
	if len(series) == 0 || span == 0 {
		p = 1
		if span > len(series) {
			err = fmt.Errorf("Bad span")
		}
		return
	}
	if span <= 0 || span > len(series) {
		return 0, fmt.Errorf("Bad span")
	}
	digits := make([]byte, span)
	l := len(series)
	temp := 1
	for x := 0; x < l; x++ {
		// fmt.Println(x, span, l)
		if x+span <= l {
			digits = []byte(series[x : x+span])
			temp, err = product(digits)
			if err != nil {
				return
			}
		}

		if temp > p {
			p = temp
		}
	}
	return
}

func product(digits []byte) (p int, err error) {
	p = 1
	for _, x := range digits {
		n, ok := strconv.Atoi(string(x))
		if ok != nil {
			return 0, fmt.Errorf("Bad input")
		}
		p *= n
	}
	// fmt.Println(digits, p)
	return
}
