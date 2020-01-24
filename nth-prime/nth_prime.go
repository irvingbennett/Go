//Package prime returns the nth prime
package prime

import (
	"github.com/jbarham/primegen.go"
)

// Nth returns the nth prime
func Nth(n int) (nPrime int, ok bool) {
	primes := []int{}
	if n == 1 {
		return 2, true
	}
	p := primegen.New()
	for i := 0; i <= n-2; i++ {
		primes = append(primes, int(p.Next()))
		ok = true
	}
	nPrime = int(p.Next())
	// fmt.Println(n, nPrime)
	// Used package primegen for this job
	// not re-inventing the wheel

	return
}
