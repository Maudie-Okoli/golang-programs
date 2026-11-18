package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	if len(a) < 2 {
		return
	}
	for c := 0; c < len(a)-1; c++ {
		for b := c + 1; b < len(a); b++ {
			if f(a[c], a[b]) > 0 {
				a[c], a[b] = a[b], a[c]
			}
		}
	}
}
