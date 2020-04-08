// Package listops is a collection of list operations
package listops

// IntList is an array of integers
type IntList []int

// Append appends one list to another
func (l IntList) Append(i IntList) {
	var m = make(IntList, len(l)+len(i))
	for x := 0; x < len(l); x++ {
		m[x] = l[x]
	}
	y := len(l)
	for x := 0; x < len(i); x++ {
		m[y] = i[x]
		y++
	}

	l = m
	return
}

// Foldr returns a list folded to the right
func (l IntList) Foldr(a func(b, c int) int) {
	for i := 0; i < len(l); i++ {
		c += a(l[i], c)
	}

	return
}
