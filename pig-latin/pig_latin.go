// Package piglatin ...
package piglatin

import (
	"fmt"
	"strings"
)

// Sentence changes to pig-latin
func Sentence(input string) (output string) {
	o := []string{}
	// fmt.Println(input, input[1:], input[0:1])
	words := strings.Fields(input)
	for _, w := range words {
		o = append(o, word(w))
	}
	output = strings.Join(o, " ")
	return
}

func word(input string) (output string) {
	switch {
	case input == "my":
		output = "ymay"

	case strings.IndexAny(input[0:1], "xy") >= 0 && strings.IndexAny(input[1:2], "aeiou") >= 0:
		output = fmt.Sprintf("%s%say", input[1:], input[0:1])

	case strings.IndexAny(input[0:1], "aeiouyx") >= 0:
		output = fmt.Sprintf("%say", input)

	case strings.Index(input[0:3], "squ") >= 0 || strings.Index(input[0:3], "thr") >= 0 || strings.Index(input[0:3], "sch") >= 0:
		output = fmt.Sprintf("%s%say", input[3:], input[0:3])

	case strings.Index(input[0:2], "ch") >= 0 || strings.Index(input[0:2], "qu") >= 0 ||
		strings.Index(input[0:2], "th") >= 0 || strings.Index(input[0:2], "rh") >= 0:
		output = fmt.Sprintf("%s%say", input[2:], input[0:2])

	default:
		output = fmt.Sprintf("%s%say", input[1:], input[0:1])
	}

	return
}
