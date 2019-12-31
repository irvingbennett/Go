package hamming

import "errors"

//Distance calculates the hamming distance
func Distance(a, b string) (int, error) {
	// Calculates the hamming distance

	if len(a) != len(b) {
		return 0, errors.New("Length does not match")
	}
	distance := 0
	if len(a) == 0 {
		return distance, nil
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
