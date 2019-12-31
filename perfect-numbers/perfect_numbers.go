// Package perfect does...
package perfect

import (
	"fmt"
	// "math"
)

var ErrOnlyPositive error = fmt.Errorf("Only positive")

/*
classifications := []struct {
		class Classification
		name  string
	}{
		{ClassificationAbundant, "ClassificationAbundant"},
		{ClassificationDeficient, "ClassificationDeficient"},
		{ClassificationPerfect, "ClassificationPerfect"},
	}
*/

//Classification is the class of a number
type Classification int

const (
	// ClassificationAbundant the sum of factors is more than then number
	ClassificationAbundant = 1
	// ClassificationPerfect the sum es equal to the number
	ClassificationPerfect = 0
	// ClassificationDeficient the sum is less
	ClassificationDeficient = -1
)

// Classify returns the classification of a number
func Classify(n int64) (class Classification, err error) {
	// fmt.Println("Got:", n)
	if n <= 0 {
		return 0, ErrOnlyPositive
	}
	factors := make([]int, 0)
	for x := 1; x < int(n); x++ {
		if n%int64(x) == 0 {
			factors = append(factors, x)
		}
	}
	sum := 0
	for _, x := range factors {
		sum += x
	}
	// fmt.Println("Factors:", factors, "Sum:", sum)
	switch {
	case int(n)-sum == ClassificationPerfect:
		class = ClassificationPerfect
		err = nil
	case int(n)-sum > 0:
		class = ClassificationDeficient
		err = nil
	case int(n)-sum < 0:
		class = ClassificationAbundant
		err = nil
	}
	return
}
