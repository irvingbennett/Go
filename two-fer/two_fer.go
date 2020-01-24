// Package twofer should have a package comment that summarizes what it's about.
package twofer

import (
	"strings"
)

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	if len(name) > 0 {
		a := []string{
			"One for ",
			name,
			", one for me.",
		}
		return strings.Join(a, "")
	}
	return "One for you, one for me."
}
