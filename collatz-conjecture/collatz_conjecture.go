package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	count := 0

	if n <= 0 {
		return count, errors.New("number is negative")
	}

	for n > 1 {
		divisible := n % 2 == 0 // check divisibilty of the number
		if divisible {
			n /= 2
		} else {
			n = 3 * n +1
		}
		count++
	}
	return count, nil
}
