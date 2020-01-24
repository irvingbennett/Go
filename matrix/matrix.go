// Package matrix returns a Matrix from a string
package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix is a slice of slice int
type Matrix [][]int

// New returns a new Matrix from a string
func New(in string) (Matrix, error) {
	lines := strings.Split(in, "\n")
	m := make(Matrix, len(lines))
	r := 0
	for _, l := range lines {
		cols := strings.Split(strings.Trim(l, " "), " ")
		// fmt.Println(cols, len(cols))
		m[r] = make([]int, len(cols))

		for c, n := range cols {
			if n == "" {
				continue
			}
			x, err := strconv.Atoi(n)
			if err != nil {
				return nil, fmt.Errorf("Bad number")
				// continue
			}
			m.Set(r, c, x)
		}
		r++
	}
	// fmt.Println(m)
	r = len(m)
	// c := len(m[r-1])
	l1 := len(m[0])
	// fmt.Println("Rows", r, "Cols", c)
	for x := 0; x < r; x++ {

		if l1 != len(m[x]) {

			return m, fmt.Errorf("Non-square matrix")
		}
	}

	return m, nil
}

// Rows returns the rows of a Matrix
func (m Matrix) Rows() [][]int {
	// fmt.Println("Rows")
	// fmt.Println(m)
	rows := len(m)
	cols := len(m[rows-1])
	mCopy := make([][]int, rows)
	for r := 0; r < rows; r++ {
		mCopy[r] = make([]int, cols)
		for c := 0; c < cols; c++ {
			mCopy[r][c] = m[r][c]
		}
	}
	// fmt.Println(mCopy)
	return mCopy
}

// Cols returns the cols of Matrix
func (m Matrix) Cols() [][]int {
	// fmt.Println("Columns")
	if len(m) == 0 {
		return *new([][]int)
	}
	rows := len(m)
	cols := len(m[rows-1])
	oldRows := rows
	oldCols := cols
	// println(rows, cols)
	// fmt.Println(m, len(m), len(m[0]))
	mCopy := make([][]int, cols)
	clm := 0
	for c := 0; c < oldCols; c++ {
		// fmt.Println(c)
		mCopy[c] = make([]int, rows)
		for r := 0; r < oldRows; r++ {
			// fmt.Println(c, r, r, clm, m[r][clm])
			mCopy[c][r] = m[r][clm]
		}
		clm++
	}
	// fmt.Println(mCopy)
	return mCopy

}

// Set puts the value in a row, col
func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 {
		return false
	}
	if row >= len(m) || col >= len(m[0]) {
		return false
	}
	m[row][col] = val
	return true
}
