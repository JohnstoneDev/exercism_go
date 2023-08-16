package anagram

import (
	_ "fmt"
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var anagrams []string

	subjectChars := strings.Split(strings.ToLower(subject), "")
	sort.Strings(subjectChars)
	sortedSubject := strings.Join(subjectChars, "")

	for _, word := range candidates {
		wordChars := strings.Split(strings.ToLower(word), "")
		sort.Strings(wordChars)
		sortedWord := strings.Join(wordChars, "")

		if subject == word || strings.EqualFold(subject, word) {
			continue
		}

		if sortedSubject == sortedWord {
			anagrams = append(anagrams, word)
		}
	}

	return anagrams
}
