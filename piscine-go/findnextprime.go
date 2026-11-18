package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		nb = 2
	}
	for {
		if isPrime(nb) {
			return nb
		}
		nb++
	}
}

func isPrime(n int) bool {
	if n < 2 || n%2 == 0 && n != 2 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
