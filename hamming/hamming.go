package hamming

import "errors"

func Distance(a, b string) (int, error) {
	count, sliceA, sliceB := 0, []byte(a), []byte(b)

	if len(a) != len(b) {
		return count, errors.New("string lengths do not match")
	}

	for i := 0; i < len(sliceA); i++{
		if sliceA[i] != sliceB[i] {
			count ++
		}
	}
	return count, nil
}
