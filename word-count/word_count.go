// Package wordcount counts words
package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is a word count map
type Frequency map[string]int

// WordCount (tt.input) counts words
func WordCount(input string) (f Frequency) {
	freqs := map[string]int{}
	wordPat := regexp.MustCompile("[a-zA-Z0-9]+(['][a-zA-Z0-9]+)?")
	for _, word := range wordPat.FindAllString(input, -1) {
		word = strings.ToLower(word)
		freqs[word]++
	}
	f = freqs
	return
}
