package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency is
func ConcurrentFrequency(s []string) FreqMap {
	m := make(chan FreqMap)
	f := FreqMap{}
	n := FreqMap{}
	for _, doc := range s {
		go func(s string) {
			n := FreqMap{}
			for _, r := range s {
				n[r]++
			}
			m <- n
		}(doc)
		n = <-m
		for x, y := range n {
			f[x] += y
		}
	}
	return f
}
