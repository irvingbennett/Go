package smallest

import (
	"fmt"
)

// Factors return prime factors
func Factors(n int64) (res []int64) {
	fmt.Println(n)
	res = make([]int64, 0)
	if n == 0 || n == 1 {
		return
	}
	if n == 2 || n == 3 {
		res = append(res, n)
		return
	}

	var number, quotient, remainder, i int64
	number = n
	fmt.Println("Number:", number)
	for i = 2; i < n; {
		prime := i
		quotient, remainder = divmod(number, int64(prime))
		// fmt.Println("Prime:", prime, "quotient:", quotient, "remainder:", remainder, "n:", n)
		if remainder == 0 {
			res = append(res, int64(prime))
			number = quotient
		} else {
			i++
		}
		if quotient == 0 {
			break
		}
	}
	fmt.Println("Prime factors:", res)
	return res
}

func divmod(numerator, denominator int64) (quotient, remainder int64) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}
