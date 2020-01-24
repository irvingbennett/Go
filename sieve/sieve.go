// Package sieve does...
package sieve

import (
	"fmt"
	"math"
)

// Sieve returns all the primes from 2 up to a given number.
func Sieve(n int) (primes []int) {
	primes = make([]int, n+16)
	fmt.Println(n)
	iter := primesOdds(uint(n + 1))
	x := 0
	for v := iter(); v <= uint(n); v = iter() {
		p := int(v)
		// fmt.Printf("%d, ", p)
		if p > 0 {
			primes[x] = p
			x++
		}
	}
	primes = primes[:x]
	return
}

func primesOdds(top uint) func() uint {
	topndx := int((top - 3) / 2)
	topsqrtndx := (int(math.Sqrt(float64(top))) - 3) / 2
	cmpsts := make([]uint, top)
	// fmt.Println(top, topndx, topsqrtndx, (topndx/32)+1)
	for i := 0; i <= topsqrtndx; i++ {
		if cmpsts[i>>5]&(uint(1)<<(uint(i)&0x1F)) == 0 {
			p := (i << 1) + 3
			for j := (p*p - 3) >> 1; j <= topndx; j += p {
				cmpsts[j>>5] |= 1 << (uint(j) & 0x1F)
			}
		}
	}
	i := -1
	return func() uint {
		oi := i
		if i <= topndx {
			i++
		}
		for i <= topndx && cmpsts[i>>5]&(1<<(uint(i)&0x1F)) != 0 {
			i++
		}
		if oi < 0 {
			return 2
		} else {
			return (uint(oi) << 1) + 3
		}
	}
}
