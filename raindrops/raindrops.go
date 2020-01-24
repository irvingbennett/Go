// Package raindrops return the sound of a number
package raindrops

import (
	"fmt"
	"strings"
)

// Convert returns a string with the sound of input
func Convert(n int) string {
	// fmt.Printf("Number: %d \n\n", n)
	sound := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}
	result := []string{"", "", ""}
	i := 0
	factors := [3]int{3, 5, 7}
	for x := range factors {
		// fmt.Printf("%d -> %s\n", factors[x], sound[factors[x]])
		if n%factors[x] == 0 {
			result[i] = sound[factors[x]]
			i++
		}
	}
	if i > 0 {
		return strings.Join(result, "")
	}
	return fmt.Sprintf("%d", n)
}
