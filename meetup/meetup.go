// Package meetup calculates the date of a meetup
package meetup

import (
	"fmt"
	"time"
)

// WeekSchedule has the day of the
type WeekSchedule int

const (
	// First day of the week
	First WeekSchedule = 1 + iota
	// Second day
	Second
	// Third day
	Third
	// Fourth day
	Fourth
	// Last day
	Last
	// Teenth day of the week
	Teenth
)

// Day returns the day of the meetup date
func Day(w WeekSchedule, day time.Weekday, month time.Month, year int) int {
	fmt.Println("Day:", w, "Weekday:", day, "Month:", month, "Year:", year, "WeekSchedule:", First)
	var dom int
	switch {
	case w == Teenth:
		fmt.Println("Teenth")
		dom = 13 // is WeekSchedule, 6 is Teenth

	case w == First:
		fmt.Println("First")
		dom = 1 // is WeekSchedule, 1 is first

	case w == Second:
		fmt.Println("Second")
		dom = 8 // is WeekSchedule, 2 is Second

	case w == Third:
		fmt.Println("Third")
		dom = 15 // is WeekSchedule, 3 is Third

	case w == Fourth:
		fmt.Println("Fourth")
		dom = 22 // is WeekSchedule, 4 is Fourth

	case w == Last:
		fmt.Println("Last")
		if m31(month) {
			dom = 25 // is WeekSchedule, 5 Last
		} else if m30(month) {
			dom = 24 // is WeekSchedule, 5 is Last
		} else {
			dom = 28 // 29 days in this month
			fmt.Println("It's February")
			for x := 0; x < 7; x++ {
				d := time.Date(year, month, dom-x, 12, 30, 0, 0, time.UTC)
				fmt.Println(d.Weekday(), d.Day())
				if d.Weekday() == day {
					return d.Day()
				}
			}
		}
	}

	for x := 0; x < 7; x++ {
		d := time.Date(year, month, dom+x, 12, 30, 0, 0, time.UTC)
		// fmt.Println(d.Weekday())
		if d.Weekday() == day {
			return d.Day()
		}
	}

	return 0
}

func m31(m time.Month) bool {
	m31 := []time.Month{time.January, time.March, time.May, time.July, time.August, time.October, time.December}
	for _, x := range m31 {
		if m == x {
			return true
		}
	}
	return false
}

func m30(m time.Month) bool {
	m31 := []time.Month{time.April, time.June, time.September, time.November}
	for _, x := range m31 {
		if m == x {
			return true
		}
	}
	return false
}
