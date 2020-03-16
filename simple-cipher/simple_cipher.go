// Package cipher provides simple cipher functions
package cipher

import (
	"bytes"
)

/*
type Cipher interface {
	Encode(string) string
	Decode(string) string
}
*/

var rng = byte('z' - 'a' + 1)

// NewCaesar returns a Caesar cipher, based on Shift cipher
func NewCaesar() Cipher {
	return NewShift(3)
}

// NewShift returns a Shift cipher, based on Vigenere cipher
func NewShift(k int) Cipher {
	if k == 0 || k > 25 || k < -25 {
		return nil
	}
	n := byte(k)
	return NewVigenere(string('a' + (rng+n)%rng))
}

// Vigenere cipher
type Vigenere struct {
	key string
}

// NewVigenere returns a new cipher
func NewVigenere(key string) Cipher {
	min := byte('a')
	max := byte('a')
	for i := 0; i < len(key); i++ {
		if key[i] < min {
			min = key[i]
		}
		if key[i] > max {
			max = key[i]
		}
	}
	if min < 'a' || max > 'z' || max == 'a' {
		return nil
	}
	return &Vigenere{key}
}

// Encode cleans a message and encodes it using the Vigenere cipher
func (cipher *Vigenere) Encode(msg string) string {
	var input bytes.Buffer
	for _, c := range msg {
		if c >= 'A' && c <= 'Z' {
			input.WriteByte(byte(c) - 'A' + 'a')
		} else if c >= 'a' && c <= 'z' {
			input.WriteByte(byte(c))
		}
	}

	clean := input.String()
	var buffer bytes.Buffer
	for i := 0; i < len(clean); i++ {
		c := clean[i]
		k := cipher.key[i%len(cipher.key)]
		s := 'a' + ((c-'a')+(k-'a'))%rng
		buffer.WriteByte(s)
	}
	return buffer.String()
}

// Decode decodes msg
func (cipher *Vigenere) Decode(msg string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(msg); i++ {
		c := msg[i]
		k := cipher.key[i%len(cipher.key)]
		s := 'a' + ((c-'a')-(k-'a')+rng)%rng
		buffer.WriteByte(s)
	}
	return buffer.String()
}
