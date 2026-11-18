package piscine

func AtoiBase(s string, base string) int {
	if !validBase(base) {
		return 0
	}

	result := 0
	b := len(base)

	for i := 0; i < len(s); i++ {
		idx := indexInBase(s[i], base)
		if idx == -1 {
			return 0
		}
		result = result*b + idx
	}

	return result
}

func validBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

func indexInBase(c byte, base string) int {
	for i := 0; i < len(base); i++ {
		if base[i] == c {
			return i
		}
	}
	return -1
}
