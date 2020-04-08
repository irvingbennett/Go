// Package smallest multiple looks for the smallest number
// that is a multiple of a range of numbers.
package smallest

import (
	"fmt"
)

// Smallest multiple looks for the smallest multiple in a range
func Smallest(r int64) (smallest int64) {
	if r == 1 || r == 2 {
		return r
	}
	for i := int64(2); i < r; i++ {
		if i == 2 {
			smallest = i * (i + 1) / GCD(i, i+1)
		} else {
			smallest = smallest * (i + 1) / GCD(smallest, i+1)
		}
	}
	fmt.Println("Range:", r, "Smallest:", smallest)
	return
}

func check(r, n int64) bool {
	for i := int64(1); i <= r; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}

func largest(rs []int64) int64 {
	product := int64(1)
	for _, x := range rs {
		product *= x
	}
	return product
}

// Find x in slice ns and return its index or -1 if not found
func Find(ns []int64, x int64) int {
	for i, n := range ns {
		if x == n {
			return i
		}
	}
	return -1
}

// GCD finds greatest common denominator
func GCD(x, y int64) int64 {
	if y == 0 {
		return x
	}
	return GCD(y, x%y)
}
