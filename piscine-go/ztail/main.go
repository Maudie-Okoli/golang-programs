package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 4 || os.Args[1] != "-c" {
		fmt.Printf("Usage: %s -c N file...\n", os.Args[0])
		os.Exit(1)
	}

	var n int
	if _, err := fmt.Sscanf(os.Args[2], "%d", &n); err != nil || n <= 0 {
		fmt.Printf("Invalid number of bytes: %s\n", os.Args[2])
		os.Exit(1)
	}

	files := os.Args[3:]
	exitCode := 0
	printedSomething := false

	for _, fname := range files {
		data, err := os.ReadFile(fname)
		if err != nil {
			fmt.Printf("%v\n", err)
			exitCode = 1
			printedSomething = true
			continue
		}

		// Print header only for successfully read files, with separation.
		if len(files) > 1 {
			if printedSomething {
				fmt.Printf("\n")
			}
			fmt.Printf("==> %s <==\n", fname)
		}

		if n >= len(data) {
			fmt.Printf("%s", data)
		} else {
			fmt.Printf("%s", data[len(data)-n:])
		}

		printedSomething = true
	}

	os.Exit(exitCode)
}
