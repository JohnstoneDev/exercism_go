package flatten

func Flatten(nested interface{}) []interface{} {
	if nested == nil {
		return []interface{}{}
	}

	flattened := []interface{}{}

	// loop over the element type
	switch elem := nested.(type) {
	case []interface{} :
		// look for interfaces nested within
		for _, item := range elem {
			if item != nil {
				flattened = append(flattened, Flatten(item)...) // recursively flatten the flattened slice
			}
		}
	default:
		if elem != nil {
			flattened = append(flattened, elem)
		}
	}

	return flattened
}