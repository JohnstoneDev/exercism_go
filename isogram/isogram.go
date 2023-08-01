package isogram

import (
	"unicode"
)

func IsIsogram(word string) bool {
	seen := make(map[rune]bool) // make a map of bools accessed using runes

	// loop through the word
	for _, char := range word {
		lowercase := unicode.ToLower(char) // change the runes to lowercase

		// only check for letters of the alphabet
		if unicode.IsLetter(lowercase) {
			// check if the character is already seen
			if seen[lowercase] {
				return false
			}
			// mark the character as seen
			seen[lowercase] = true
		}
	}

	return true
}
