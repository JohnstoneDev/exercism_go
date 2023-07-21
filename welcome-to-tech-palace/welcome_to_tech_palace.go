package techpalace

import (
	s "strings"
	"fmt"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	customerName := s.ToUpper(customer)
	formattedString := fmt.Sprintf("Welcome to the Tech Palace, %s", customerName)
	return formattedString
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starLine := s.Repeat("*", numStarsPerLine)
	decoratedString := fmt.Sprintf("%s\n%s\n%s", starLine, welcomeMsg, starLine)
	return decoratedString
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return s.TrimSpace(s.ReplaceAll(oldMsg, "*", ""))
}