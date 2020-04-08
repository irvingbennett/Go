// Package dominoes ...
package dominoes

import (
	"fmt"
)

// Domino type
type Domino [2]int

func reverse(d Domino) Domino {
	return Domino{d[1], d[0]}
}

// MakeChain returns a chain of dominoes
func MakeChain(dominoes []Domino) (d []Domino, ok bool) {
	fmt.Println(dominoes)
	fmt.Println("------------------------------")
	fmt.Println()

	if len(dominoes) == 0 {
		return dominoes, true
	}
	copy := dominoes
	d, ok = chain(copy)
	fmt.Println("------------------------------")
	fmt.Println()
	return
}

func check(dominoes []Domino) bool {

	if dominoes[0][0] == dominoes[len(dominoes)-1][1] {
		return true
	}
	return false
}

func remove(slice []Domino, i int) []Domino {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func chain(d []Domino) ([]Domino, bool) {
	if len(d) == 1 {
		return d, check(d)
	}
	start := d
	dominoes := len(d)
	tries := 0
	for {
		try := []Domino{}

		try = append(try, start[tries])
		start = remove(start, tries)
		dominoes--
		i := 0
		for len(start) > 0 {
			fmt.Println(try)
			if try[len(try)-1][1] == start[i][0] {
				try = append(try, start[i])
				start = remove(start, i)
				dominoes--
			} else {
				i++
			}
			if i >= len(start) {
				i = 0
			}
		}
		if check(try) && len(start) == 0 {
			fmt.Println(try, true)
			return try, true
		}
		tries++
		start = d
		if tries == len(d) {
			fmt.Println(try, false)
			return try, false
		}
	}
}
