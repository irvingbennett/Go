// Package anagram returns the anagrama from a list
package anagram

import (
	"fmt"
	"strings"
)

// Detect returns the anagram of a word from a list
func Detect(s string, list []string) (a []string) {
	fmt.Println(s, list)
	for _, c := range list {
		if anagram(strings.ToLower(s), strings.ToLower(c)) {
			// fmt.Println(a, c)
			a = append(a, c)
		}
	}

	return
}

func anagram(w, c string) bool {
	if w == c {
		return false
	}
	lw := len(w)
	lc := len(c)
	if lw != lc {
		return false
	}
	seenW := make(map[byte]int, lw)
	seenC := make(map[byte]int, lw)
	for i := range w {
		seenW[w[i]]++
		seenC[c[i]]++
	}
	// fmt.Println(seenW, seenC)
	for wl, wi := range seenW {
		if seenC[wl] != wi {
			return false
		}
	}
	return true
}
