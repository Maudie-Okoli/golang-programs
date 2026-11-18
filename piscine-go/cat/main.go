package main

import (
	"bufio"
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]

	// If no arguments, read from stdin
	if len(args) == 0 {
		reader := bufio.NewReader(os.Stdin)
		for {
			line, err := reader.ReadString('\n')
			if len(line) > 0 {
				printString(line)
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				printString("ERROR: " + err.Error() + "\n")
				os.Exit(1)
			}
		}
		return
	}

	// If arguments are provided, treat each as a file
	for _, fileName := range args {
		file, err := os.Open(fileName)
		if err != nil {
			printString("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}
		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			if len(line) > 0 {
				printString(line)
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				printString("ERROR: " + err.Error() + "\n")
				os.Exit(1)
			}
		}
	}
}
