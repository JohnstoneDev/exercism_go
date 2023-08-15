package resistorcolor

// Colors should return the list of all colors.
func Colors() []string {
	colors := []string{
		"black",
		"brown",
		"red",
		"orange",
		"yellow",
		"green",
		"blue",
		"violet",
		"grey",
		"white",
	}

	return colors
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	colors := Colors();

	for i, value := range colors {
		if value == color {
			return i
		}
	}

	return -1
}
