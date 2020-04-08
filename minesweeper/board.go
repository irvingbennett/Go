package minesweeper

import "bytes"

// Board is the type for a board filled with mines
type Board [][]byte

func (b Board) String() string {
	return "\n" + string(bytes.Join(b, []byte{'\n'}))
}
