package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// calculates the weigh of a number provided as a string
func calculateWeight(str string) int {
  sum := 0
  integers := strings.Split(str, "")

  for _, numStr := range integers {
    value, _ := strconv.Atoi(numStr)
    sum += value
  }

  return sum
}

// custom sorting type & methods
type ByWeight []string

func (numStr ByWeight) Len() int {
  return len(numStr)
}

func (numStr ByWeight) Swap(i, j int) {
  numStr[i], numStr[j] = numStr[j], numStr[i]
}

func (numStr ByWeight) Less(i, j int) bool {
  weighI  := calculateWeight(numStr[i])
  weightJ := calculateWeight(numStr[j])

  // if the weight is the same, return the first one in the list
  if weighI == weightJ {
    return numStr[i] < numStr[j]
  }

  return weighI < weightJ
}

func OrderWeight(str string) string {
	numbers := strings.Fields(str) // Convert input string to a slice of strings
	sort.Sort(ByWeight(numbers))   // Sort using the custom sorting type

	return strings.Join(numbers, " ") // Join the sorted slice back to a string
}

func main() {
  fmt.Println(OrderWeight("103 123 4444 99 2000"))
  fmt.Println(OrderWeight("2000 103 123 4444 99"))
}
