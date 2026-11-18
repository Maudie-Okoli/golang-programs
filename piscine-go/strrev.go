package piscine

func StrRev(s string) string {
	r := []rune(s)
	i := 0
	j := len(r) - 1

	for i < j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}

	return string(r)
}
