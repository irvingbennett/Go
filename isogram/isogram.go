package isogram

import "strings"

// IsIsogram returns true if a word is an isogram
func IsIsogram(s string) bool {
	word := strings.ToLower(s)
	letters := make(map[rune]int)

	for i := 0; i < len(word); i++ {
		letter := rune(word[i])
		letters[letter]++
		if letter != rune('-') && letter != rune(' ') {
			if letters[rune(word[i])] > 1 {
				return false
			}
		}
	}

	return true

}
