// Package rectangles ...
package rectangles

// Count counts rectangles
func Count(shape []string) int {
	if len(shape) == 0 {
		return 0
	}
	if len(shape) == 1 {
		return 0
	}
	rectangles := 0
	corner := '+'
	l := len(shape)
	w := len(shape[0])
	points := make([][]rune, l)
	for i := 0; i < l; i++ {
		l := []rune(shape[i])
		points[i] = l

	}
	for x := 0; x < l-1; x++ {
		for y := 0; y < w-1; y++ {
			for x2 := 1; x2 < l; x2++ {
				if points[x][y] != corner {
					continue
				}
				for y2 := 1; y2 < w; y2++ {
					if x == x2 || y == y2 {
						continue
					}
					if points[x][y] == corner && points[x][y2] == corner &&
						points[x2][y] == corner && points[x2][y2] == corner {
						if check(points, x, y, x2, y2) {
							rectangles++
						}
					}
				}
			}
		}
	}

	return rectangles

}

func check(sqr [][]rune, i, l, i2, l2 int) bool {
	for x := i; x <= i2; x++ {
		for y := l; y <= l2; y++ {
			if x == i || x == i2 {
				if sqr[x][y] != '+' && sqr[x][y] != '-' && sqr[x][y] != '|' {
					return false
				}
			} else if sqr[x][l] != '+' && sqr[x][l] != '|' ||
				sqr[x][l2] != '+' && sqr[x][l2] != '|' {
				return false
			}
		}

	}
	if l < l2 && i < i2 {
		return true
	}
	return false

}
