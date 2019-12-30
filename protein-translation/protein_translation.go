// Package protein does...
package protein

import (
	"fmt"
)

/*

Codon                 | Protein
:---                  | :---
AUG                   | Methionine
UUU, UUC              | Phenylalanine
UUA, UUG              | Leucine
UCU, UCC, UCA, UCG    | Serine
UAU, UAC              | Tyrosine
UGU, UGC              | Cysteine
UGG                   | Tryptophan
UAA, UAG, UGA         | STOP

*/

// ErrStop is stop error
var ErrStop error = fmt.Errorf("Stop error")

// ErrInvalidBase is invalid base
var ErrInvalidBase = fmt.Errorf("Invalid base")

// FromCodon (input) returns...
func FromCodon(input string) (output string, err error) {
	var c string
	codon := make(map[int]string, 0)
	codon[1] = "Methionine"
	codon[2] = "Phenylalanine"
	codon[3] = "Leucine"
	codon[4] = "Serine"
	codon[5] = "Tyrosine"
	codon[6] = "Cysteine"
	codon[7] = "Tryptophan"
	codon[8] = "STOP"

	RNA := make(map[string]int, 0)
	RNA["AUG"] = 1
	RNA["UUU"] = 2
	RNA["UUC"] = 2
	RNA["UUA"] = 3
	RNA["UUG"] = 3
	RNA["UCU"] = 4
	RNA["UCC"] = 4
	RNA["UCA"] = 4
	RNA["UCG"] = 4
	RNA["UAU"] = 5
	RNA["UAC"] = 5
	RNA["UGU"] = 6
	RNA["UGC"] = 6
	RNA["UGG"] = 7
	RNA["UAA"] = 8
	RNA["UAG"] = 8
	RNA["UGA"] = 8

	l := 3
	// fmt.Println("Input: ", input, "Output:", output)
	for x := 0; x <= len(input)-3; x += 3 {
		c = input[x:l]
		// fmt.Println(codon[RNA[c]])
	}
	output, ok := codon[RNA[c]]
	// fmt.Println("Output:", output, ok, err)
	if len(output) == 0 || !ok {
		// fmt.Println("Invalid base:", output, ok, err)
		err = ErrInvalidBase
	}

	if output == "STOP" {
		output = ""
		err = ErrStop
	}
	return
}

// FromRNA (input)
func FromRNA(input string) (output []string, err error) {
	// fmt.Println("Input: ", input, "Output:", output)
	l, idx := 3, 0
	output = make([]string, 32)
	for x := 0; x <= len(input)-3; x += 3 {
		c := input[x : x+l]
		protein, err := FromCodon(c)
		if err != nil {
			if err == ErrStop {
				break
			}
			output = output[:idx]
			// fmt.Println(output, err)
			return output, err
		}
		output[idx] = protein
		// fmt.Println(output)
		idx++
	}
	output = output[:idx]
	// fmt.Println(output)
	return
}
