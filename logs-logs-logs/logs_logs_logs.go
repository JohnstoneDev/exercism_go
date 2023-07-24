package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	recommendation, search, weather := '❗', '🔍', '☀'

	for _, char := range log {
		if char == recommendation {
			return "recommendation"
		}

		if char == search {
			return "search"
		}

		if char == weather {
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runes := []rune(log)

	for i, char := range runes {
		if char == oldRune {
			runes[i] = newRune
		}
	}
	return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
