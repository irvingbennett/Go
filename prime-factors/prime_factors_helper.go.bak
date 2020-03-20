// Package prime return prime factors of a number
package prime

// Generate produces a sequence on integers.
// Send the sequence 2, 3, 4, ... to channel 'out'
func Generate(out chan<- int) {
	for i := 2; ; i++ {
		out <- i // Send 'i' to channel 'out'
	}
}

// Filter removes primes.
// Copy the values from 'in' channel to 'out' channel,
// removing the multiples of 'prime' by counting.
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

// Sieve filters primes
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
