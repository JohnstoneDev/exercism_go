package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime (layers []string, layerPrepTime int) int {
	if layerPrepTime == 0 {
		return 2 * len(layers)
	}
	return len(layers) * layerPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities (layers []string) (noodles int, sauce float64) {
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		}

		if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient (list, ingredients []string) {
	lastIngredient := list[len(list) - 1]
	ingredients[len(ingredients) - 1] = lastIngredient
}

// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe (quantities []float64, portions int) []float64 {
	portionsRequired := make([]float64, len(quantities))

	for i, quantity := range quantities {
		portionsRequired[i] = quantity * float64(portions) / 2
	}
	return portionsRequired
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
