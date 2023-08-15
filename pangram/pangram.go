package pangram

import "strings"

func IsPangram(input string) bool {
	alphabetMap := make(map[rune]int) // make a map to store the alphabet
	upcased := strings.ToUpper(input) // transform the input to uppercase

	// populate alphabet map with capital letters, initialize wit zero
	for r := 'A'; r <= 'Z'; r++ {
		alphabetMap[r] = 0
	}

	// loop through upcased counting all the letters, increment them in alphabetMap
	for _, letter := range upcased {
		_, found := alphabetMap[letter]

		if found {
			alphabetMap[letter]++
		}
	}

	// Check to see that there are no missing letters & return
	for _, num := range alphabetMap {
		if num == 0 {
			return false
		}
	}

	return true
}
