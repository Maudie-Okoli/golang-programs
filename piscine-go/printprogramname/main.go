package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]

	// remove the path before the program name
	lastSlash := 0
	for i := 0; i < len(name); i++ {
		if name[i] == '/' {
			lastSlash = i + 1
		}
	}

	name = name[lastSlash:]

	for _, r := range name {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
