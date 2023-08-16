// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
	"regexp"
)

func shouting(sentence string) bool {
	// regex to pick letters only
	pattern := regexp.MustCompile(`[a-zA-Z]`)
	onlyLetters := pattern.FindAllString(sentence, -1)

	lettersSentence := strings.Join(onlyLetters, "")

	// make a slice of the original sentence's letters only
	original := []rune(lettersSentence)
	var uppercase []rune

	for _, letter := range sentence {
		if unicode.IsUpper(letter) {
			uppercase = append(uppercase, letter)
		}
	}

	return len(original) == len(uppercase)
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimRight(remark, " "); // remove white space at the end of the remark

	question := strings.HasSuffix(remark, "?")
	yell := shouting(remark)
	yellQuestion := yell && question

	noWords := regexp.MustCompile(`[a-zA-Z]`).MatchString(remark)
	silence := !noWords

	if silence {
		if question {
			return "Sure."
		}
		// Cheekc for numbers
		numbers := regexp.MustCompile(`[0-9]`).MatchString(remark)
		if numbers {
			return "Whatever."
		}

		return "Fine. Be that way!"
	}

	if yellQuestion {
		return "Calm down, I know what I'm doing!"
	}

	if yell {
		return "Whoa, chill out!"
	}

	if question {
		return "Sure."
	}

	return "Whatever."
}
