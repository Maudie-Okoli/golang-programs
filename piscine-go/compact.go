package piscine

func Compact(ptr *[]string) int {
	// Create a new slice to store non-empty strings
	var result []string

	// Loop through original slice
	for _, str := range *ptr {
		if str != "" {
			result = append(result, str)
		}
	}

	// Replace original slice with the compacted one
	*ptr = result

	// Return number of non-empty strings
	return len(result)
}
