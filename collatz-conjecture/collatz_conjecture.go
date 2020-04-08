// Package collatzconjecture returns the number of steps to reach 1
package collatzconjecture

import (
	"fmt"
)

// CollatzConjecture (input int) (steps int, err error)
func CollatzConjecture(input int) (steps int, err error) {
	if input == 0 || input < 0 {
		err = fmt.Errorf("improper input")
		return
	}

	var n = float64(input)
	steps = 0
	fmt.Println(input)
	for n != 1 {
		if int(n)%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		steps++
	}

	return steps, nil
}
