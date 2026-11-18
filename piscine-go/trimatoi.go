package piscine

func TrimAtoi(s string) int {
	sign := 1
	num := 0
	started := false

	for _, c := range s {
		if c == '-' && !started {
			sign = -1
			started = true
		}
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
			started = true
		}
	}

	return sign * num
}
