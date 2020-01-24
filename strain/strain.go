// Package strain filters a set
package strain

// Ints is a list of integers
type Ints []int

// Lists is a list of list integer
type Lists [][]int

// Strings is a slice of strings
type Strings []string

//
// Then complete the exercise by implementing these methods:
//
//    (Ints) Keep(func(int) bool) Ints

// Keep keeps those items that pass test
func (i Ints) Keep(f func(int) bool) (k Ints) {
	if len(i) == 0 {
		return
	}
	k = make(Ints, 0)
	for _, n := range i {
		if f(n) {
			k = append(k, n)
		}
	}
	return k
}

// Discard discards items that don't pass the filter function
func (i Ints) Discard(f func(int) bool) (k Ints) {
	if len(i) == 0 {
		return
	}
	k = make(Ints, 0)
	for _, n := range i {
		if !f(n) {
			k = append(k, n)
		}
	}
	return k

}

// Keep for lists
func (i Lists) Keep(f func([]int) bool) (k Lists) {
	k = make(Lists, 0)
	for _, n := range i {
		if f(n) {
			k = append(k, n)
		}
	}
	return k
}

// Keep for strings
func (i Strings) Keep(f func(string) bool) (k Strings) {

	k = make(Strings, 0)
	for _, n := range i {
		if f(n) {
			k = append(k, n)
		}
	}
	return k
}
