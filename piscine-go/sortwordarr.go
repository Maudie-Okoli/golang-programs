package piscine

func SortWordArr(a []string) {
	if len(a) < 2 {
		return
	}
	for t := 0; t < len(a); t++ {
		for s := t + 1; s < len(a); s++ {
			if a[t] > a[s] {
				a[t], a[s] = a[s], a[t]
			}
		}
	}
}
