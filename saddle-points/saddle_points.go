// Package matrix returns a list of saddle points
package matrix

// Pair is an int matrix type
type Pair [2]int

// Saddle returns the saddle points of a matrix
func (m Matrix) Saddle() (pairs []Pair) {
	// fmt.Println(m)
	for x := 0; x < len(m); x++ {
		largest := 0
		p := Pair{-1, -1}
		for y := 0; y < len(m[x]); y++ {
			if m[x][y] > largest {
				largest = m[x][y]
				p[0] = x
				p[1] = y
			}
		}
		s := m[int(p[0])][int(p[1])]
		for r := 0; r < len(m); r++ {
			if s > m[r][int(p[1])] {
				p = Pair{-1, -1}
				break
			}
		}
		if p[0] >= 0 && p[1] >= 0 {
			pairs = append(pairs, p)
		}

	}
	return
}
