// Package twelve outputs the twelve days of Christmas song
package twelve

import (
	"fmt"
)

type verse struct {
	day   string
	verse string
}

// Verses holds the lines of the song
var Verses = []verse{
	{"first", "a Partridge in a Pear Tree"},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

// Verse composes the verse for one day
func Verse(day int) string {
	v := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", Verses[day-1].day)
	// fmt.Println("Day: ", day)
	for d := day - 1; d >= 1; d-- {
		v += fmt.Sprintf("%s, ", Verses[d].verse)
	}
	if day == 1 {
		v += fmt.Sprintf("%s.", Verses[0].verse)
	} else {
		v += fmt.Sprintf("and %s.", Verses[0].verse)
	}
	/*
		if day != 12 {
			v += "\n"
		} else {
			v += ""
		}
	*/
	return v
}

// Song puts the verses together
func Song() string {
	var s string
	for d := 1; d <= 12; d++ {
		s += Verse(d)
		if d != 12 {
			s += "\n"
		}
	}
	return s
}
