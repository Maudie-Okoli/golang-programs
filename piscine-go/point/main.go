package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)

	printStr("x = ")

	// Print points.x inline (no helper)
	n := points.x
	p := 1
	for t := n; t >= 10; t = t / 10 {
		p = p * 10
	}

	for p > 0 {
		d := n / p // digit as int 0..9

		// build rune digit without conversions
		ch := '0'
		for i := 0; i < d; i++ {
			ch++
		}
		z01.PrintRune(ch)

		n = n % p
		p = p / 10
	}

	printStr(", y = ")

	// print points.y inline (no helper)
	n2 := points.y
	p2 := 1
	for t := n2; t >= 10; t = t / 10 {
		p2 *= 10
	}
	for p2 > 0 {
		d := n2 / p2

		ch := '0'
		for i := 0; i < d; i++ {
			ch++
		}
		z01.PrintRune(ch)

		n2 = n2 % p2
		p2 = p2 / 10
	}

	z01.PrintRune('\n')
}
