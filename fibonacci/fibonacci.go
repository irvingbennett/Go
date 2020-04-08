// Package fibonacci returns a sequence of Fibonacci numbers
package fibonacci

// Fibonacci return a slice of Fibonacci number up to number
func Fibonacci(n int) (f int) {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		f = Fibonacci(n-2) + Fibonacci(n-1)
	}
	return
}
