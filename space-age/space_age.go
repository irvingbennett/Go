// Package space calculates the age of a person
// on each planet
package space

// Planet type is used for orbital peiods
type Planet string

// SECONDS is the time, in seconds,
// that it takes for an orbit of the Earth
// around the Sun.
const SECONDS = 31557600

// Age calculates de age of a person on a planet
func Age(s float64, p Planet) float64 {

	period := map[string]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	// fmt.Printf("Seconds: %f Planet: %v  \n\n", s, p)
	return s / (SECONDS * period[string(p)])
}
