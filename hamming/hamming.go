package hamming

import "errors"

func Distance(a, b string) (int, error) {
	count, sliceA, sliceB := 0, []byte(a), []byte(b) // create slices containing the characters in a & b as ASCII bytes

	// Compare string lengths & return 0 & an error if they do not match
	if len(a) != len(b) {
		return count, errors.New("string lengths do not match")
	}

	// loop through the slices comparing the values at every index, increment count
	// for bytes that do not match
	for i := 0; i < len(sliceA); i++{
		if sliceA[i] != sliceB[i] {
			count ++
		}
	}
	return count, nil
}
