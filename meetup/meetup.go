// Package meetup calculates the date of a meetup
package meetup

import (
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
	// fmt.Println("Day:", w, "Weekday:", day, "Month:", month, "Year:", year, "WeekSchedule:", First)
	var dom int
	switch {
	case w == Teenth:
		dom = 13 // is WeekSchedule, 6 is Teenth

	case w == First:
		dom = 1 // is WeekSchedule, 1 is first

	case w == Second:
		dom = 8 // is WeekSchedule, 2 is Second

	case w == Third:
		dom = 15 // is WeekSchedule, 3 is Third

	case w == Fourth:
		dom = 22 // is WeekSchedule, 4 is Fourth

	case w == Last:
		if m31(month) {
			dom = 25 // is WeekSchedule, 5 Last
		} else if m30(month) {
			dom = 24 // is WeekSchedule, 5 is Last
		} else {
			dom = 22 // 28 days in this month

			if IsLeapYear(year) {
				dom = 23 // 29 days in this month
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

// IsLeapYear returns true if a year is leap
func IsLeapYear(l int) bool {
	if l%4 == 0 {
		if l%100 == 0 {
			if l%400 != 0 {
				return false
			}
		}
		return true
	}
	return false
}
