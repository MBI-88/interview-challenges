package main

import (
	"flag"
	"fmt"
)

func fillBoard(b [][]string) [][]string {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			b[i][j] = "."
		}
	}
	return b
}

func isSafe(b [][]string, row, col int) bool {
	for c := 0; c < len(b); c++ {
		if b[row][c] == "Q" {
			return false
		}
	}

	for r := 0; r < len(b); r++ {
		if b[r][col] == "Q" {
			return false
		}
	}

	c := col
	r := row
	for c < len(b) && r < len(b) {
		if b[r][c] == "Q" {
			return false
		}
		c++
		r++
	}

	c = col
	r = row
	for c < len(b) && r >= 0 {
		if b[r][c] == "Q" {
			return false
		}
		c++
		r--
	}

	c = col
	r = row
	for c >= 0 && r >= 0 {
		if b[r][c] == "Q" {
			return false
		}
		c--
		r--
	}

	c = col
	r = row
	for c >= 0 && r < len(b) {
		if b[r][c] == "Q" {
			return false
		}
		c--
		r++
	}

	return true
}

func showBoard(b [][]string) {
	for row := 0; row < len(b); row++ {
		for col := 0; col < len(b); col++ {
			fmt.Printf("%s", b[row][col])
		}
		fmt.Println()
	}
}

func nQueens(queens, dm int) {
	board := make([][]string, dm)
	for i := range board {
		board[i] = make([]string, dm)
	}
	board = fillBoard(board)

	for q := 0; q < queens; q++ {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board); j++ {
				if isSafe(board, i, j) {
					board[i][j] = "Q"
				}
			}
		}

	}

	showBoard(board)
}

var (
	queens *int 
	dm *int
)

func init() {
	queens = flag.Int("queens", 0, "total queens")
	dm = flag.Int("dm", 0, "board dimentions")
}

func main() {
	flag.Parse()
	
	nQueens(*queens, *dm)
}
