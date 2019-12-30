// Package allergies returns allergens from a score
package allergies

/*
* eggs (1)
* peanuts (2)
* shellfish (4)
* strawberries (8)
* tomatoes (16)
* chocolate (32)
* pollen (64)
* cats (128)
 */

var allergens = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

var food = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

// AllergicTo (test.score, response.substance)
func AllergicTo(score uint, substance string) (is bool) {

	if score&allergens[substance] == allergens[substance] {
		return true
	}

	return
}

// Allergies returns...
func Allergies(score uint) (allergies []string) {
	// fmt.Println(score)
	allergies = make([]string, 0)
	for _, s := range food {
		i := allergens[s]
		if score&i == i {
			allergies = append(allergies, s)
		}
	}

	return
}
