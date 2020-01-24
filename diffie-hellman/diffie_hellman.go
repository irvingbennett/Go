// Package diffiehellman does...
package diffiehellman

import (
	"math"
	"math/big"
	"math/rand"
	"time"
)

// Diffie-Hellman-Merkle key exchange
//
// Step 1:   PrivateKey(p *big.Int) *big.Int
// Step 2:   PublicKey(private, p *big.Int, g int64) *big.Int
// Step 2.1: NewPair(p *big.Int, g int64) (private, public *big.Int)
// Step 3:   SecretKey(private1, public2, p *big.Int) *big.Int
//
// Private keys should be generated randomly.

// PrivateKey (p *big.Int) *big.Int
func PrivateKey(p *big.Int) (k *big.Int) {
	// fmt.Println(p)
	var s *rand.Rand
	key := new(big.Int)
	limit := new(big.Int).Sub(p, big.NewInt(2))
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	if &s == nil {
		s = seed
	} else {
		seed = rand.New(rand.NewSource(int64(math.Round(math.Pi * 10000000))))
		s = nil
	}
	// fmt.Printf("Type: %T\n", seed)
	k = key.Rand(seed, limit).Add(key, big.NewInt(2))
	return
}

// PublicKey (private, p *big.Int, g int64) *big.Int
func PublicKey(private, p *big.Int, g int64) *big.Int {
	// fmt.Println(private)
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

// NewPair (p *big.Int, g int64) (private, public *big.Int)
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	// fmt.Println(p, g)
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

// SecretKey (private1, public2, p *big.Int) *big.Int
func SecretKey(private1, public2, p *big.Int) *big.Int {
	// fmt.Println(private1, public2)

	return new(big.Int).Exp(public2, private1, p)
}
