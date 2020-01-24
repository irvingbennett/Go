// Package wordy parses words
package wordy

import (
	"strconv"
	"strings"
)

// Answer return the answer to the question
func Answer(q string) (answer int, ok bool) {
	// fmt.Println(q)
	if strings.Index(q, "What is ") < 0 {
		ok = false
		return
	}
	exp := q[len("What is "):]
	// fmt.Println(exp)
	operands := split(exp)
	var num int
	operator := ""
	lastOperator := ""
	var lastNumber bool
	for i := 0; i < len(operands); i++ {

		if n, err := strconv.Atoi(operands[i]); err != nil {
			ok = false
			if operands[i] == "by" {
				continue
			}
			if lastOperator != "" {
				ok = false
				return
			}
			lastOperator = operands[i]
			lastNumber = false

			switch {
			case i == 0:
				ok = false
				return
			case operands[i] == "plus":
				operator = "+"
			case operands[i] == "minus":
				operator = "-"
			case operands[i] == "multiplied":
				operator = "*"
			case operands[i] == "divided":
				operator = "/"
			case operands[i] == "cubed":
				ok = false
				return
			}

		} else {
			switch {
			case lastNumber:
				ok = false
				return

			case operator == "" && lastOperator == "":
				num = n

			case operator == "+":
				num += n

			case operator == "-":
				num -= n

			case operator == "*":
				num *= n

			case operator == "/":
				num /= n
			}
			lastOperator = ""
			lastNumber = true
			ok = true
		}

	}
	answer = num

	return
}

func split(q string) []string {
	o := strings.Split(q, " ")
	if strings.Index(o[len(o)-1], "?") > 0 {
		o[len(o)-1] = o[len(o)-1][:strings.Index(o[len(o)-1], "?")]
	}
	// fmt.Println(o)
	return o
}
