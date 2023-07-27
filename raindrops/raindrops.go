package raindrops

import "strconv"

func Convert(number int) string {
	factor3 := number % 3 == 0
	factor5 := number % 5 == 0
	factor7 := number % 7 == 0

	if factor3 && factor5 && factor7 {
		return "PlingPlangPlong"
	}

	if factor3 && factor5 {
		return "PlingPlang"
	}

	if factor5 && factor7 {
		return "PlangPlong"
	}

	if factor3 && factor7 {
		return "PlingPlong"
	}

	if factor3 {
		return "Pling"
	}

	if factor5 {
		return "Plang"
	}

	if factor7 {
		return "Plong"
	}


	return strconv.Itoa(number)
}
