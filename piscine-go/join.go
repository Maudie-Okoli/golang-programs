package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}

	result := strs[0] // start with the first element
	for p := 1; p < len(strs); p++ {
		result += sep + strs[p] // add seperator before each next element
	}
	return result
}
