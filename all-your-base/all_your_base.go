// Package allyourbase transforms numbers from one base to another
package allyourbase

import (
	"fmt"
	"math"
)

// output, err := ConvertToBase(c.inputBase, c.inputDigits, c.outputBase)

// ConvertToBase does the conversion
func ConvertToBase(base int, input []int, outputBase int) (output []int, err error) {
	fmt.Printf("Input base: %d, input: %d, output base: %d \n", base, input, outputBase)
	var decimal int
	for x := 0; x < len(input); x++ {
		decimal += input[x] * int(math.Pow(float64(base), float64(x)))
	}
	fmt.Println("Output:", decimal)

	for x := 1; decimal > 0; x++ {
		colVal := int(math.Pow(float64(outputBase), float64(x)))
		r := decimal % colVal
		decimal -= r
		var digit = 0
		if r > 0 {
			digit = r / colVal
		}
		fmt.Println("colVal", colVal, "decimal", decimal, "digit", digit, "r", r)
		output = append(output, digit)

		if output[0] == 0 {
			output = output[1:]
		}
	}

	return
}
