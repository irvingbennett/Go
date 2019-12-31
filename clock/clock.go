// Package clock implements a clock
package clock

import (
	"fmt"
)

// Clock is a struct for keeping time
type Clock struct {
	//thisTime time.Time
	Hour    int
	Minutes int
}

// New creates a new clock
func New(t, m int) Clock {
	var c Clock
	c.Hour = t % 24
	if m >= 60 {
		c.Hour += int(m / 60)
		c.Hour = c.Hour % 24
	}
	c.Minutes = m % 60
	if m < 0 {
		if m <= -60 {
			c.Hour += int(m/60) - 1
			if m%60 == 0 {
				c.Hour++
			}
			c.Hour = c.Hour % 24
			if c.Hour < 0 {
				c.Hour += 24
			}
		} else {
			c.Hour--
		}
		if c.Hour < 0 {
			c.Hour = c.Hour % 24 * -1
		}
		if c.Minutes < 0 {

			c.Minutes += 60
		}
	}
	if c.Hour < 0 {
		c.Hour += 24
	}
	// n := time.Now()
	// year, month, day := n.Date()
	// c.thisTime = time.Date(year, month, day, t, m, 0, 0, time.UTC)
	// fmt.Printf("New: %d, %d, Got hour %d, minute %d\n \n", t, m, c.Hour, c.Minutes)

	return c
}

// String returns a string representation of the time
func (c Clock) String() string {
	// layout := "04:05"
	// s := c.thisTime.Format(layout)
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minutes)
}

// Add adds a duration to the time
func (c Clock) Add(m int) Clock {
	// duration := fmt.Sprintf("%dm", m)
	hours := int(m / 60)
	minutes := m % 60
	c.Hour += hours
	c.Minutes += minutes
	if c.Minutes >= 60 {
		c.Minutes -= 60
		c.Hour++
	}
	if c.Hour >= 24 {
		c.Hour -= 24
	}
	c.Hour = c.Hour % 24
	// fmt.Printf("Add: %d, Had hour %d, minute %d, Got hours: %d, minutes %d\n \n", m, c.Hour, c.Minutes, hours, minutes)
	// d, _ := time.ParseDuration(duration)
	// c.thisTime.Add(d)
	return c
}

// Subtract takes duration off the time
func (c Clock) Subtract(m int) Clock {

	// duration := fmt.Sprintf("%dm", m)
	hours := int(m / 60)
	minutes := m % 60
	c.Hour -= hours
	if minutes > c.Minutes {
		c.Hour--
		c.Minutes = c.Minutes - minutes + 60
	} else {
		c.Minutes -= minutes
	}
	if c.Minutes < 0 {
		c.Hour--
		c.Minutes = 60 - c.Minutes
	}
	if c.Hour < 0 {
		c.Hour = c.Hour % 24
		c.Hour += 24
		if c.Hour == 24 {
			c.Hour = 0
		}
	}
	// fmt.Printf("Sub: %d, Had hour %d, minute %d, Got hours: %d, minutes %d\n \n", m, c.Hour, c.Minutes, hours, minutes)

	// d, _ := time.ParseDuration(duration)
	// c.thisTime.Add(-d)
	return c
}
