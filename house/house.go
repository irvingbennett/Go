// Package house tells the verses of the house that jack built
package house

import (
	"fmt"
)

type verse struct {
	sub string
	did string
}

var verses []verse

// Song returns the full song
func Song() string {
	s := ""
	for i := 1; i <= 12; i++ {
		if i == 12 {
			s = s + Verse(i)
		} else {
			s = s + Verse(i) + "\n\n"
		}
	}
	// fmt.Println(s)
	return s
}

// Verse return a verse from the song
func Verse(i int) (v string) {

	verses := make([]verse, 24)
	verses = []verse{
		{sub: "house", did: "Jack built"},
		{sub: "malt", did: "lay"},
		{sub: "rat", did: "ate"},
		{sub: "cat", did: "killed"},
		{sub: "dog", did: "worried"},
		{sub: "cow with the crumpled horn", did: "tossed"},
		{sub: "maiden all forlorn", did: "milked"},
		{sub: "man all tattered and torn", did: "kissed"},
		{sub: "priest all shaven and shorn", did: "married"},
		{sub: "rooster that crowed in the morn", did: "woke"},
		{sub: "farmer sowing his corn", did: "kept"},
		{sub: "horse and the hound and the horn", did: "belonged to"},
	}

	if i == 1 {
		v = "This is the house that Jack built."
		// fmt.Println(i, v)
	} else {
		v = fmt.Sprintf("This is the %s\n", verses[i-1].sub)
		// fmt.Println(i, v)
		for l := i - 2; l > 0; l-- {
			v = v + fmt.Sprintf("that %s the %s\n", verses[l+1].did, verses[l].sub)
			// fmt.Println(l, v)
		}
		v = v + "that lay in the house that Jack built."
	}
	// fmt.Println(v)
	return
}
