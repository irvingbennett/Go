// Package grains calculate the debt of ric
package grains

import (
	"errors"
	"math"
)

// Square calculates grains of rice square for
// each square on a chessboard
func Square(s int) (uint64, error) {
	if s < 1 || s > 64 {
		return 0, errors.New("Invalid value")
	}
	return uint64(math.Pow(2, float64(s-1))), nil
}

// Total returns the full square
func Total() uint64 {
	var sum uint64
	for x := 1; x <= 64; x++ {
		sum += uint64(math.Pow(2, float64(x-1)))
	}
	return uint64(sum)
}
