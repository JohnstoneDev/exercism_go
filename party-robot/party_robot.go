package partyrobot

import (
	f "fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return f.Sprintf("Welcome to my party, %v!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return f.Sprintf("Happy birthday %v! You are now %v years old!", name, age)
}

// Format table number to append zeroes
func formatTableNumber(number int) string {
	return f.Sprintf("%03d", number)
}
// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var string1, string2, string3 string;
	string1 = f.Sprintf("%s\n", Welcome(name))
	string2 = f.Sprintf("You have been assigned to table %v. Your table is %v, exactly %.1f meters from here.\n", formatTableNumber(table), direction, distance)
	string3 = f.Sprintf("You will be sitting next to %v.", neighbor)

	return string1 + string2 + string3
}