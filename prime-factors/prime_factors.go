// Package prime return prime factors of a number
package prime

import (
	"fmt"

	"math"
)

// Factors return prime factors
func Factors(n int64) (res []int64) {
	fmt.Println(n, math.MaxInt64)
	if n > math.MaxInt64 {
		return
	}
	res = make([]int64, 0)
	if n == 0 || n == 1 {
		return
	}

	sv := make(chan int) // Create a new channel
	go Sieve(sv)

	var number, quotient, remainder, i int64
	number = n
	for i = 0; i < n; {
		prime := <-sv
		quotient, remainder = divmod(number, int64(prime))
		fmt.Println("Prime:", prime, "q:", quotient, "r:", remainder, "n:", n)
		if remainder == 0 {
			res = append(res, int64(prime))
			number = quotient
		} else {
			prime := <-sv
			i = int64(prime)
		}
	}
	return res
}

func divmod(numerator, denominator int64) (quotient, remainder int64) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}

// Send the sequence 2, 3, 4, ... to channel 'out'
func Generate(out chan<- int) {
	for i := 2; ; i++ {
		out <- i // Send 'i' to channel 'out'
	}
}

// Copy the values from 'in' channel to 'out' channel,
//   removing the multiples of 'prime' by counting.
// 'in' is assumed to send increasing numbers
func Filter(in <-chan int, out chan<- int, prime int) {
	m := prime + prime // first multiple of prime
	for {
		i := <-in // Receive value from 'in'
		for i > m {
			m = m + prime // next multiple of prime
		}
		if i < m {
			out <- i // Send 'i' to 'out'
		}
	}
}

// The prime sieve: Daisy-chain Filter processes
func Sieve(out chan<- int) {
	gen := make(chan int) // Create a new channel
	go Generate(gen)      // Launch Generate goroutine
	for {
		prime := <-gen
		out <- prime
		ft := make(chan int)
		go Filter(gen, ft, prime)
		gen = ft
	}
}
