package wordy

import (
	"regexp"
	"strconv"
)

func Answer(question string) (int, bool) {
	reNumbers := regexp.MustCompile(`-?\d+`) //extract numbers and operations from string
	reOperations := regexp.MustCompile(`(plus|minus|multiplied by|divided by)`)

	// find all numbers and operations
	numbers := reNumbers.FindAllString(question, -1)
	operations := reOperations.FindAllString(question, -1)

	// convert the first number to an integer
	number, err := strconv.Atoi(numbers[0])

	if err != nil {
		return 0, false
	}

	// if there are no operators return number & false
	if len(operations) == 0 {
		return number, false
	}

	// iterate through remaining numbers & operations
	for i, num := range numbers[1:]{
		numString, err := strconv.Atoi(num)

		if err != nil {
			return number, false
		}

		switch(operations[i-1]) {
		default :
			return 0, false
		case "plus" :
			number += numString
		case "minus" :
			number -= numString
		case "multiplied by" :
			number *= numString
		case "divided by" :
			if numString != 0 {
				number /= numString
			} else {
				return 0, false
			}
		}
	}

	return number, true
}
