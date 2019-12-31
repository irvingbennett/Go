// Package triangle tests a triangle and returns its kind
package triangle

import (
	"math"
	"sort"
)

// Kind of a triangle
type Kind string

const (
	// NaT is not a triangle
	NaT = "Nat" // not a triangle
	// Equ is equal
	Equ = "Equ" // equilateral
	// Iso is isoceles
	Iso = "Iso" // isosceles
	// Sca is scalene
	Sca = "Sca" // scalene
)

// KindFromSides returns the kind of a triangle
func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	pinf := math.Inf(1)
	ninf := math.Inf(-1)
	for _, n := range sides {
		if math.IsNaN(n) || n == pinf || n == ninf {
			return Kind(NaT)
		}
	}
	sort.Float64s(sides)
	// fmt.Println(sides)
	var k Kind
	switch {
	case a <= 0 || b <= 0 || c <= 0:
		k = NaT

	case sides[0]+sides[1] < sides[2]:
		k = NaT

	case a == b && b == c:
		k = Equ

	case sides[0] == sides[1] || sides[2] == sides[1]:
		k = Iso

	default:
		k = Sca
	}

	return k
}
