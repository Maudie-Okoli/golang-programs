package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20 { // prevent overflow
		return 0
	}
	if nb == 0 || nb == 1 { // base cases
		return 1
	}
	return nb * RecursiveFactorial(nb-1) // recursive call
}
