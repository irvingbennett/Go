// Package brackets verifies all brackets match
package brackets

import (
	"fmt"
	"strings"
)

type match struct {
	pair    string
	index   int
	matched bool
}

// Bracket returns true if all brackets match
func Bracket(s string) (correct bool) {
	if s == "" {
		return true
	}
	pairs := map[string]string{"{": "}", "[": "]", "(": ")"}
	matches := make([]match, 0)
	var itsCorrect bool
	// var m match
	correct = false
	fmt.Println(s)
	first := strings.IndexAny(s, "{[(")
	// fmt.Println(first, string(s[first]))
	if first >= 0 { // look for match
    match := string(s[first])
    next := strings.Index(s[match+1:], match) 
    if next > 1 {
      last := strings.Index(s, pairs[string(s[first])])
      matchPair := last
      if last > 0 {
        last := strings.Index(s[matchPair+1:], pairs[string(s[first])])
        if last < 0 {
          return false
        }
      }
    }
    fmt.Println("Match:",match)
		last := strings.Index(s, pairs[string(s[first])])
		// fmt.Println(last)
		if last > 0 {

			itsCorrect = true
			if last+1 == len(s) {
				return true
			} else if last-first > 1 {
				itsCorrect = Bracket(s[first : last+1])
			}
			if itsCorrect && last+1 < len(s) {
				return Bracket(s[last+1:])
			}
		}

	}
	fmt.Println("Its", itsCorrect, matches)
	correct = itsCorrect
	return itsCorrect
}