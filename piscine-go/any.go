package piscine

func Any(f func(string) bool, a []string) bool {
	for _, h := range a {
		if f(h) {
			return true
		}
	}
	return false
}
