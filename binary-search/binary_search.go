// Package binarysearch return an item
package binarysearch

// SearchInts searches a list for a value
func SearchInts(l []int, v int) (p int) {
	if len(l) == 0 {
		return -1
	}
	b, t := 0, len(l)
	p = len(l) / 2
	pb, pt := b, t
	for l[p] != v {
		// fmt.Printf("Bottom: %d, Top: %d, p: %d\n", b, t, p)
		if l[p] > v {
			t = p
			p = (t - b) / 2
		} else {
			b = p
			p = (t-b)/2 + b
		}
		if pb == b && pt == t {
			break
		}
		pb, pt = b, t
		// fmt.Printf("Bottom: %d, Top: %d, p: %d\n", b, t, p)
	}
	if l[p] != v {
		p = -1
	}
	return p
}

// SearchIntsStd returns the position of an item in a list
func SearchIntsStd(a []int, value int) int {
	low := 0
	high := len(a) - 1
	for low <= high {
		mid := (low + high) / 2
		if a[mid] > value {
			high = mid - 1
		} else if a[mid] < value {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
