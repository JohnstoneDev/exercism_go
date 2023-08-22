package main

import (
	"fmt"
)

func Permutations(str string) []string {
	perms := []string{} // permutations
	var generate func(string, string) // function that will generate string
	seen := make(map[string]bool) // map that stores a generated permutation

	// use generate function recursively to create permutations
	generate = func(remaining, current string){
		if len(remaining) == 0 {
			if !seen[current] {
				perms = append(perms, current)
				seen[current] = true
			}
			return
		}

		for i, char := range remaining {
			remainingWithoutChar := remaining[:i] + remaining[i + 1:]
			generate(remainingWithoutChar, current+string(char))
		}
	}

	generate(str, "") // start generating from the input with an empty stting as the base case

	return perms
}

func main() {
	fmt.Println(Permutations("a"), "<= Expecting [a]")
	fmt.Println(Permutations("ab"), "<= Expecting [ab, ba]")
}