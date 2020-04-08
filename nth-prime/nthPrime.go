//Package prime returns the nth prime
package main

import (
	"fmt"

	"github.com/jbarham/primegen.go"
)

// Nth returns the nth prime
func Nth(n int) (nPrime int, ok bool) {
	primes := 0
	if n == 1 {
		return 2, true
	}
	p := primegen.New()
	for i := 0; i < n-2; i++ {
		primes += int(p.Next())
	}
	ok = true
	nPrime = int(p.Next())
	fmt.Println(n, nPrime, primes)
	// Used package primegen for this job
	// not re-inventing the wheel

	return
}

// SumPrimes adds all primes below a value
func SumPrimes(n int) (sum int, ok bool) {
	if n == 1 {
		return 2, true
	}
	p := primegen.New()
	var next int
	for {
		next = int(p.Next())
		if next < n {
			sum += next
		} else {
			break
		}
	}
	ok = true
	nPrime := next
	fmt.Println(n, nPrime, sum)
	// Used package primegen for this job
	// not re-inventing the wheel

	return
}

func main() {
	// nth, ok := Nth(2000000)
	sum, ok := SumPrimes(2000000)
	if !ok {
		fmt.Println("error...")
	}
	// fmt.Println("2,000,000 prime is", nth)
	fmt.Println("Sum of primes below 2,000,000 is", sum)
}
