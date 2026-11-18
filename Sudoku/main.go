package main

import (
	"fmt"
	"os"
)

const N = 9

func isValid(grid [][]rune, row, col int, ch rune) bool {
	for i := 0; i < N; i++ {
		if grid[row][i] == ch || grid[i][col] == ch {
			return false
		}
	}
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[startRow+i][startCol+j] == ch {
				return false
			}
		}
	}
	return true
}

func solve(grid [][]rune) bool {
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if grid[row][col] == '.' {
				for ch := '1'; ch <= '9'; ch++ {
					if isValid(grid, row, col, ch) {
						grid[row][col] = ch
						if solve(grid) {
							return true
						}
						grid[row][col] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	grid := make([][]rune, N)
	for i := 0; i < N; i++ {
		if len(os.Args[i+1]) != 9 {
			fmt.Println("Error")
			return
		}
		grid[i] = []rune(os.Args[i+1])
		for _, v := range grid[i] {
			if !(v == '.' || (v >= '1' && v <= '9')) {
				fmt.Println("Error")
				return
			}
		}
	}

	if !solve(grid) {
		fmt.Println("Error")
		return
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%c", grid[i][j])
			if j != N-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
