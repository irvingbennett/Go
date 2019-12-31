package sublist

// Relation is the return value of Sublist
type Relation string

func isSublist(a, b []int) bool {
	lenA, lenB := len(a), len(b)
	for i := 0; i <= lenB-lenA; i++ {
		if areEqual(a, b[i:i+lenA]) {
			return true
		}
	}
	return false
}

func areEqual(a, b []int) bool {
	lenA, lenB := len(a), len(b)
	if lenA != lenB {
		return false
	}
	for i := 0; i < lenA; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Sublist check if a is a sublist of b
func Sublist(a, b []int) Relation {
	lenA, lenB := len(a), len(b)
	switch {
	case lenA == lenB:
		switch areEqual(a, b) {
		case true:
			return Relation("equal")
		case false:
			return Relation("unequal")
		}
	case lenA > lenB:
		switch isSublist(b, a) {
		case true:
			return Relation("superlist")
		case false:
			return Relation("unequal")
		}
	case lenA < lenB:
		switch isSublist(a, b) {
		case true:
			return Relation("sublist")
		case false:
			return Relation("unequal")
		}
	}

	return Relation("unequal")
}
