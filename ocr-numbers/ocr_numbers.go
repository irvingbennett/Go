// Package ocr recognizes numbers
package ocr

import (
	"strings"
)

var digits map[string]string

func splitIn(o []string) (out []string) {
	// o := strings.Split(in, "\n")
	o = o[1:4]
	// fmt.Println(len(o))
	// fmt.Println(o)
	l := len(o[1]) / 3
	/*
		  fmt.Println("Len:", l)
				for i, s := range o{
				  fmt.Println(i, "Len:", len(s), s)
				}
	*/
	for n := 0; n < l; n++ {
		s := []string{o[0][n*3 : n*3+3], o[1][n*3 : n*3+3], o[2][n*3 : n*3+3]}
		// fmt.Println(s)
		out = append(out, strings.Join(s, "\n"))
	}
	// fmt.Println("Out:", out)
	return
}

// Recognize recognizes a whole string
func Recognize(in string) (out []string) {
	lines := strings.Split(in, "\n")

	for i := 0; i < len(lines)-4; i += 4 {
		o := []string{}
		input := splitIn(lines[i : i+4])
		// fmt.Println("Input:", input, "Len:", len(input))
		for _, x := range input {
			o = append(o, recognizeDigit(x))
		}
		out = append(out, strings.Join(o, ""))
	}
	return
}

func recognizeDigit(in string) string {
	setupDigits()
	i := strings.ReplaceAll(strings.TrimSpace(in), "\n", "")
	// fmt.Println("Digit:", digits[i])
	var val string
	if v, ok := digits[i]; !ok {
		val = "?"
	} else {
		val = v
	}
	return val
}

func setupDigits() {
	if len(digits) <= 0 {
		digits = make(map[string]string, 10)
		ds := map[string]string{
			`
 _ 
| |
|_|
	 `: "0",
			`
   
  |
  |
  `: "1",
			`
 _ 
 _|
|_ 
  `: "2",
			`
 _ 
 _|
 _|
  `: "3",
			`
   
|_|
  |
  `: "4",
			`
 _ 
|_ 
 _|
  `: "5",
			`
 _ 
|_ 
|_|
  `: "6",
			`
 _ 
  |
  |
  `: "7",
			`
 _ 
|_|
|_|
  `: "8",
			`
 _ 
|_|
 _|
  `: "9",
		}
		for t, n := range ds {
			d := strings.ReplaceAll(strings.TrimSpace(t), "\n", "")
			// d = strings.ReplaceAll(d, " ", "")
			digits[d] = n
			// fmt.Println("Digits:", n, "Linear:", d, "Len:", len(d))
			// fmt.Println()
		}
	}
}
