// Package leap returns true if a year is leap
package leap

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
