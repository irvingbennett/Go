// Package pascal computes pascal's triangle for a given row
package pascal

// Triangle computes pascal's triangle for a given row
func Triangle(n int) (t [][]int) {
	t = make([][]int, n+1)
	for i := 0; i < n; i++ {
		t[i] = make([]int, i+1)
		if i == 0 {
			t[i][0] = 1
		} else if i == 1 {
			t[i][0] = 1
			t[i][1] = 1
		} else {
			for x := 0; x <= i; x++ {
				if x == 0 {
					t[i][x] = 1
				} else if x == i {
					t[i][x] = 1
				} else {
					t[i][x] = t[i-1][x-1] + t[i-1][x]
				}
			}
		}
	}
	t = t[:n]
	return
}
