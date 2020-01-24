// Package beer recites the beer song
package beer

import (
	"fmt"
)

// Verse returns one verse of the song
func Verse(n int) (verse string, err error) {
	if n > 99 || n < 0 {
		err = fmt.Errorf("invalid vers")
		return
	}
	// fmt.Println(n)
	if n == 0 {
		verse = "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
		return
	}
	if n == 1 {
		verse = "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"
		return
	}
	if n == 2 {
		verse = "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n"
		return
	}
	verse = fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1)
	return
}

// Verses return a range of verses
func Verses(l, u int) (verses string, err error) {
	// fmt.Println("Lower:", l, "Upper:", u)
	if l < u {
		err = fmt.Errorf("lower bound is larger that upper bound")
		return
	}
	if l > 99 {
		err = fmt.Errorf("invalid start")
		return
	}

	if u < 0 {
		err = fmt.Errorf("invalid stop")
		return
	}
	for i := l; i >= u; i = i - 1 {
		v, _ := Verse(i)
		verses = verses + v + "\n"
	}
	return
}

// Song returns the entire song
func Song() (song string) {
	song, _ = Verses(99, 0)
	return
}
