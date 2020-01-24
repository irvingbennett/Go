// Package series returns series within a string
package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) (series []string) {
	if n > len(s) {
		return
	}
	for x := 0; x <= len(s)-n; x++ {
		series = append(series, s[x:n+x])
	}
	return
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		return ""
	}
	return s[0:n]
}

// First returns the first substring of s with length n.
func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}
	return s[0:n], true
}
