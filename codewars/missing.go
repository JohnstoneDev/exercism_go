package main

import (
	"fmt"
)

// finds the missing alphabetical letter in a slice of letters
func FindMissingLetter(chars []rune) rune {
	var alphabetRange []rune // range of letters from the first index to the last alphabetically
	var missingLetter rune // variable to store the missing letter

	// populate the alphabet slice
	for char := chars[0]; char <= chars[len(chars) - 1]; char ++ {
		alphabetRange = append(alphabetRange, char)
	}

	// loop through the slices comparing indices
	for i, letter := range alphabetRange {
		if chars[i] != letter {
			missingLetter = letter
			break
		}
	}

	return missingLetter
}

func main(){
	fmt.Println(FindMissingLetter([]rune{'a', 'b', 'c', 'd', 'f'}))
	fmt.Println(FindMissingLetter([]rune{'O', 'Q', 'R', 'S'}))
} 