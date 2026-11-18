package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for o, p := range a {
		result[o] = f(p)
	}
	return result
}
