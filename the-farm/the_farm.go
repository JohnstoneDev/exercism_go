package thefarm

import (
	"errors"
	"fmt"
)

// Invalid Cows Error

type InvalidCowsError struct {
	cows int
	message string
}

func (e InvalidCowsError) Error() string {
	return  fmt.Sprintf("%d cows are invalid: %v", e.cows, e.message)
}

// Divides the fodder available with the number of cows passed in
func DivideFood(fod FodderCalculator, cows int) (float64, error) {
	fodder, err := fod.FodderAmount(cows)
	if err == nil {
		factor, err := fod.FatteningFactor()
		if err != nil {
			return 0, err
		}

		return fodder * factor / float64(cows), nil
	}
	return 0, err
}

// Validates the Number of cows & Divides the Fodder
func ValidateInputAndDivideFood(fod FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(fod, cows)
}

// Validates the number of cows & uses a custom error
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cows,"there are no negative cows"}
	}

	if cows == 0 {
		return &InvalidCowsError{cows,"no cows don't need food"}
	}

	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
