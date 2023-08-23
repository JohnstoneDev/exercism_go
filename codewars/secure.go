package main

import (
	"fmt"
	"unicode"
)

// validates that a user string is alphanumeric
func alphanumeric(str string) bool {
	// must have at least one character
	if len(str) <= 0 {
		return false
	}

	// loop over the string & check if characters are letters || numbers
	for _, char := range str {
		allowed := unicode.IsLetter(char) || unicode.IsNumber(char)
		if !allowed {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(alphanumeric(""), "<= Expecting false")
	fmt.Println(alphanumeric(".*?"), "<= Expecting false")
	fmt.Println(alphanumeric("hello world_"), "<= Expecting false")
	fmt.Println(alphanumeric("PassW0rd"), "<= Expecting true")
}