// Package robotname generates random robot names
package robotname

import (
	"math/rand"
	"time"
)

var names map[string]bool

// Robot is a robot type
type Robot struct {
	name string
}

// Name makes a new name for robot
func (r *Robot) Name() (string, error) {
	if names == nil {
		names = make(map[string]bool)
	}
	if len(r.name) > 0 {
		return r.name, nil
	}
	name := make([]rune, 5)
	rand.Seed(time.Now().UnixNano())
	for {
		for i := 0; i < 5; i++ {
			if i < 2 {
				name[i] = rune(int('A') + rand.Intn(27))
			} else {
				name[i] = rune(int('0') + rand.Intn(8))
			}
		}
		if !names[string(name)] {
			names[string(name)] = true
			break
		}
	}
	r.name = string(name)

	return r.name, nil
}

// Reset gives the robot a new name
func (r *Robot) Reset() (string, error) {
	name := make([]rune, 5)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		if i < 2 {
			name[i] = rune(int('A') + rand.Intn(27))
		} else {
			name[i] = rune(int('0') + rand.Intn(9))
		}
	}
	r.name = string(name)
	return r.name, nil
}
