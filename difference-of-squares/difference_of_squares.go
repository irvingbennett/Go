// Package diffsquares finds the difference of the
// square of the sums and the sum of the squares of
// the first n number
package diffsquares

import "fmt"

// SumOfSquares returns the sum of the squares
func SumOfSquares(n int) int {
	sum := 0
	for x := 1; x <= n; x++ {
		sum += x * x
	}
	return sum
}

// SquareOfSum returns the square of the sum
func SquareOfSum(n int) int {
	sum := 0
	sum = int(n * (n + 1) / 2)
	return (sum * sum)
}

// Difference returns the difference of the sums
func Difference(n int) int {
	difference := SquareOfSum(n) - SumOfSquares(n)
	fmt.Println("Difference of", n, "is", difference)
	return difference
}
