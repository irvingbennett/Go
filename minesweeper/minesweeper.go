//Package minesweeper plays a bomb of a game
package minesweeper

import (
	"fmt"
	"strings"
)

// Count places the count of adjacent mines around each mine
func (b Board) Count() (err error) {

	// fmt.Println(b)
	// fmt.Printf("%d, %d\n", len(b), len(b[0]))
	if err = check(b); err != nil {
		return err
	}

	for x := 0; x < len(b); x++ {

		for y := 0; y < len(b[x]); y++ {
			count := 0
			if b[x][y] == ' ' {
				for n := x - 1; n <= x+1; n++ {
					for m := y - 1; m <= y+1; m++ {
						if b[n][m] == '*' {
							count++
						}
					}
				}
			}

			if count > 0 {
				b[x][y] = byte('0' + count)
			}
		}

	}
	// fmt.Println(b)
	return nil
}

func check(b Board) error {
	l := len(b)
	m := len(b[0])
	top := "+" + strings.Repeat("-", m-2) + "+"
	if string(b[0]) != top ||
		string(b[len(b)-1]) != top {
		return fmt.Errorf("bad edge")
	}
	for x := 0; x < l; x++ {
		if x > 0 && x < l-1 {
			if b[x][0] != '|' {
				return fmt.Errorf("bad edge")
			}
		}

		if len(b[x]) != m {
			return fmt.Errorf("bad length")
		}
		if x > 0 && x < l-1 {
			if b[x][len(b[x])-1] != '|' {
				return fmt.Errorf("bad edge")
			}
		}
		for y := 0; y < len(b[x]); y++ {
			if !strings.ContainsRune("+|- *", rune(b[x][y])) {
				return fmt.Errorf("bad character")
			}
		}
	}

	return nil
}
