package piscine

func Capitalize(s string) string {
	r := []rune(s)
	newWord := true

	for i, c := range r {
		if isAlnum(c) {
			if newWord {
				// uppercase
				if c >= 'a' && c <= 'z' {
					r[i] = c - 32
				}
				newWord = false
			} else {
				// lowercase
				if c >= 'A' && c <= 'Z' {
					r[i] = c + 32
				}
			}
		} else {
			newWord = true
		}
	}
	return string(r)
}

func isAlnum(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}
