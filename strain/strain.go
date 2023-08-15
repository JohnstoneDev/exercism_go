package strain

import (
	_ "fmt"
)

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var newSlice []int // make a new slice

	// loop through input & filter
	for _, num := range i {
		pass := filter(num)

		if pass {
			newSlice = append(newSlice, num)
		}
	}

	return newSlice
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var newSlice []int

	for _, num := range i {
		fail := !filter(num)

		if fail {
			newSlice = append(newSlice, num)
		}
	}

	return newSlice
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var newList Lists

	for _, list := range l {
		pass := filter(list)

		if pass {
			newList = append(newList, list)
		}
	}

	return newList
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var stringSlice Strings

	for _, str := range s {
		pass := filter(str)

		if pass {
			stringSlice = append(stringSlice, str)
		}
	}

	return stringSlice
}
