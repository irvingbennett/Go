// Package darts scores the landing of darts
package darts

import (
	"math"
)

type point struct {
	x float64
	y float64
}

func (p point) radius() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

// Score returns the score of a dart on a board
func Score(x, y float64) int {
	dart := point{float64(x), float64(y)}
	radius := dart.radius()
	switch {
	case radius <= 1.0:
		return 10

	case radius <= 5.0:
		return 5

	case radius <= 10.0:
		return 1

	default:
		return 0

	}
}
