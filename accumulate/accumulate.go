// Package accumulate applies a function to a slice of strings
package accumulate

// Accumulate applies a function to a slice of strings
func Accumulate(given []string, f func(string) string) (r []string) {
	if len(given) == 0 {
		return
	}
	for _, s := range given {
		r = append(r, f(s))
	}
	return
}
