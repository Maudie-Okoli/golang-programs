package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	solve(0, &board)
}

func solve(col int, board *[8]int) {
	if col == 8 {
		printBoard(board)
		return
	}

	for row := 0; row < 8; row++ {
		if isSafe(board, col, row) {
			board[col] = row
			solve(col+1, board)
		}
	}
}

func isSafe(board *[8]int, col, row int) bool {
	for c := 0; c < col; c++ {
		r := board[c]
		if r == row || // same row
			r-c == row-col || // same diagonal
			r+c == row+col { // same diagonal
			return false
		}
	}
	return true
}

func printBoard(board *[8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(board[i] + 1 + '0'))
	}
	z01.PrintRune('\n')
}
