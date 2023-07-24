package parsinglogfiles

import (
	"regexp"
	"fmt"
)
func IsValidLine(text string) bool {
	pattern := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return pattern.MatchString(text)
}

func SplitLogLine(text string) []string {
	pattern := regexp.MustCompile(`\<[-=~*]*\>`)
	return pattern.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	pattern := regexp.MustCompile(`(?i)\".*password.*\"`)
	count := 0
	for _, matched := range lines {
		if pattern.MatchString(matched){
			count += 1
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	pattern := regexp.MustCompile(`end-of-line[0-9]*`)
	return pattern.ReplaceAllString(text, "")
}


func TagWithUserName(lines []string) []string {
	pattern := regexp.MustCompile(`User\s+(\w+)`)

	for i, line := range lines {
		matched := pattern.FindStringSubmatch(line)
		if len(matched) > 1 {
			line = fmt.Sprintf("[USR] %v %v", matched[1], line)
		}
		lines[i] = line
	}

	return lines
}
