// Package queenattack returns true if the queens car strike each other
package queenattack

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// CanQueenAttack returns true if the queens car strike each other
func CanQueenAttack(white, black string) (attack bool, err error) {
	w, b := strings.ToUpper(white), strings.ToUpper(black)

	// fmt.Println(w, b)
	if w == b {
		return false, fmt.Errorf("Same square")
	}
	if !onBoard(w, b) {
		return false, fmt.Errorf("Off the board")
	}

	if w[0] == b[0] || w[1] == b[1] {
		// same column or row
		attack = true
		err = nil
	}

	if sameDiagonal(w, b) {
		attack = true
		err = nil
	}

	return
}

func onBoard(w, b string) (on bool) {
	on = true
	if w[0] < 'A' || b[0] < 'A' || w[0] > 'H' || b[0] > 'H' || w[1] < '1' || b[1] < '1' || w[1] > '8' || b[1] > '8' {
		on = false
	}
	return
}

// isDiagonal checks if the two given positions share common left or right diagonals
func sameDiagonal(p1 string, p2 string) bool {
	// for code clarity, assign each part to variables
	p1File := rune(p1[0])
	p1Rank, _ := strconv.Atoi(string(p1[1]))
	p2File := rune(p2[0])
	p2Rank, _ := strconv.Atoi(string(p2[1]))

	// check if the absolute values of the diff of files and ranks are equal which indicates a diagonal
	return math.Abs(float64(p2File-p1File)) == math.Abs(float64(p2Rank-p1Rank))
}
