// Package pythagorean seeks numbers that make a pythagorean triplet
package pythagorean

import (
	// "fmt"
	"math"
	"sort"
)

// Triplet is a three number array of pythagoren numbers
type Triplet [3]int

// Equal is true if there is already an equal triplet
func (t Triplet) Equal(n Triplet) bool {
	for x, i := range t {
		// fmt.Println(n[x], i)
		if n[x] != i {
			return false
		}
	}
	return true
}

// Range returns a list of triplets in a range
func Range(min, max int) []Triplet {
	t := make([]Triplet, 0, 32)
	idx := 0
	for a := min; a <= max; a++ {
		a2 := a * a
		for b := min; b <= max; b++ {
			b2 := b * b
			s := float64(a2) + float64(b2)
			c := math.Sqrt(s)
			if int64(c)*int64(c) == int64(s) && int(c) <= max {
				triple := []int{a, b, int(c)}
				sort.Ints(triple)
				exists := false
				for _, trpl := range t {
					if trpl.Equal(Triplet{triple[0], triple[1], triple[2]}) {
						exists = true
						continue
					}
				}
				if !exists {
					t = append(t, Triplet{triple[0], triple[1], triple[2]})
					// fmt.Println(t)
					idx++
				}
			}

		}
	}
	return t
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	t := make([]Triplet, 0, 32)
	idx := 0
	for a := 1; a <= p; a++ {
		a2 := a * a
		for b := 1; b <= p; b++ {
			b2 := b * b
			s := float64(a2) + float64(b2)
			c := math.Sqrt(s)
			if int64(c)*int64(c) == int64(s) && int64(a)+int64(b)+int64(c) == int64(p) {
				triple := []int{a, b, int(c)}
				sort.Ints(triple)
				exists := false
				for _, trpl := range t {
					if trpl.Equal(Triplet{triple[0], triple[1], triple[2]}) {
						exists = true
						continue
					}
				}
				if !exists {
					t = append(t, Triplet{triple[0], triple[1], triple[2]})
					// fmt.Println(t)
					idx++
				}
			}

		}
	}
	return t
}
