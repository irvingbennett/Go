// Package etl extract scrabble scores
package etl

import (
	"fmt"
	"strings"
)

// Transform transforms scrabble data to another format
func Transform(m map[int][]string) (t map[string]int) {
	fmt.Println(m)
	t = make(map[string]int)
	for x, s := range m {
		fmt.Println(x, s)
		for _, l := range s {
			lower := strings.ToLower(l)
			t[lower] = x
		}
	}
	fmt.Println(t)
	return
}
