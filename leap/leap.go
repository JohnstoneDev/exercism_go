// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// function that checks whether a year is leap or not
func IsLeapYear(year int) bool {
	four , hund, fourHund := (year % 4 == 0), (year % 100 == 0), (year % 400 == 0)

	if four {
		if !hund || fourHund {
			return true
		}
	}

	return false
}
