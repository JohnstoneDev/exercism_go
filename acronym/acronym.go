package acronym

import (
	"regexp"
	"strings"
)

// takes a string & returns the abbreviation
func Abbreviate(s string) string {
	pattern := `[-_]`
	regex := regexp.MustCompile(pattern)

	matched := regex.ReplaceAllString(s, " ")
	separated := strings.Split(matched, " ");

	var firsts []string

	for _, word := range separated {
		if len(word) > 0 {
			firsts = append(firsts, string(word[0]))
		}
	}

	abbrv := strings.Join(firsts, "")

	return strings.ToUpper(abbrv)
}
