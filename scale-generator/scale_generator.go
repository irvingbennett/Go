// Package scale does...
package scale

import (
	"fmt"
)

var aScale = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}

// Scale (test.tonic, test.interval) generates a melody
func Scale(tonic, interval string) (scale []string) {
	fmt.Println("Tonic:", tonic, "Interval:", interval, "Len:", len(aScale))
	x := 0
	count := 0
	scale = make([]string, len(aScale))
	start := false
	for {
		fmt.Printf("Tonic: %s Scale:  %s \n", tonic, aScale[x])
		if tonic == aScale[x] || start {
			scale[count] = aScale[x]
			start = true
			count++
		}
		x++
		if x%12 == 0 {
			x = 0
		}
		if count == 12 {
			break
		}
	}
	fmt.Println(scale, x, count)
	return
}
