// Package stringset implements sets of strings
package stringset

import (
	"fmt"
	"strings"
)

// Set implements a set of strings
type Set map[string]bool

// New returns a new, empty set
func New() Set {
	s := make(Set, 0)
	return s
}

// NewFromSlice ([]string) Set
func NewFromSlice(members []string) (n Set) {
	n = make(Set)
	for _, m := range members {
		n[m] = true
	}
	return
}

// (s Set) String() string
func (s Set) String() (str string) {
	set := make([]string, 0)
	for m := range s {
		set = append(set, fmt.Sprintf("\"%s\"", m))
	}
	str = strings.Join(set, ", ")
	str = concat("{", strings.TrimSpace(str), "}")

	return
}

// IsEmpty () bool
func (s Set) IsEmpty() bool {
	if len(s) >= 1 {
		return false
	}
	return true
}

// Has (string) bool
func (s Set) Has(m string) bool {
	if s[m] {
		return true
	}
	return false
}

// Subset (s1, s2 Set) bool
func Subset(s1, s2 Set) bool {
	if len(s1) > len(s2) {
		return false
	}
	for m := range s1 {
		if s2[m] != true {
			return false
		}
	}
	return true
}

// Disjoint (s1, s2 Set) bool
func Disjoint(s1, s2 Set) bool {
	for m := range s1 {
		if s2[m] {
			return false
		}
	}
	return true
}

// Equal (s1, s2 Set) bool
func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	for m := range s1 {
		if s2[m] == false {
			return false
		}
	}
	return true
}

// Add (string)
func (s Set) Add(m string) {
	s[m] = true
}

// Intersection (s1, s2 Set) Set
func Intersection(s1, s2 Set) Set {
	s3 := New()
	for m := range s1 {
		if s2[m] {
			s3[m] = true
		}
	}
	return s3
}

// Difference (s1, s2 Set) Set
func Difference(s1, s2 Set) Set {
	if s1.IsEmpty() {
		return s1
	}
	s3 := New()
	for m := range s1 {
		if s2[m] == false {
			s3[m] = true
		}
	}
	/*
	for m := range s2 {
		if s1[m] == false {
			s3[m] = true
		}
	}
	*/
	return s3
}

// Union (s1, s2 Set) Set
func Union(s1, s2 Set) Set {
	s3 := New()
	for m := range s1 {
		s3[m] = true
	}
	for m := range s2 {
		s3[m] = true
	}
	return s3
}

func concat(s ...string) string {
	buf := make([]byte, 512)
	count := 0

	for _, triad := range s {
		count += copy(buf[count:], triad)
	}
	buf = buf[:count]
	// fmt.Println("Count:", count)
	return string(buf)
}
