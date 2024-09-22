package main

import "fmt"

func fillMines(moves, matriz [][]int, y, x, limit int) {
	if matriz[y][x] == 0 {
		for _, m := range moves {
			Y, X := y+m[0], x+m[1]
			if (X >= 0 && X <= limit) && (Y >= 0 && Y <= limit) {
				if matriz[Y][X] == 9 {
					matriz[y][x] += 1
				}
			}
		}
	}

}

func convert1To9(matriz [][]int) {
	for y := 0; y < len(matriz); y++ {
		for x := 0; x < len(matriz[y]); x++ {
			if matriz[y][x] == 1 {
				matriz[y][x] = 9
			}
		}
	}
}

func checkLimit(matriz [][]int) bool {
	var (
		countY int
	)
	for y := 0; y < len(matriz); y++ {
		for x := 0; x < len(matriz[y]); x++ {
			countY++
		}
	}

	if countY%len(matriz) != 0 {
		return false
	}

	return true
}

func mineSweeperDetector(matriz [][]int) [][]int {
	if !checkLimit(matriz) {
		return matriz
	}

	limit := len(matriz) - 1

	var moves = [][]int{
		{0, -1}, // (y,x)
		{1, -1},
		{-1, -1},
		{0, 1},
		{-1, 1},
		{1, 1},
		{-1, 0},
		{1, 0},
	}

	convert1To9(matriz)

	for y := 0; y < len(matriz); y++ {
		for x := 0; x < len(matriz[y]); x++ {
			fillMines(moves, matriz, y, x, limit)
		}
	}

	return matriz
}

/*
	direcctions -> (-1,0),(-1,1),(-1,-1),(1,0),(1,-1)(1,1)(0,-1),(0,1)
*/

func main() {
	m1 := [][]int{
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 1, 0, 1},
		{1, 1, 0, 0},
	}
	m2 := [][]int{
		{0, 1, 0, 0, 1},
		{0, 0, 1, 0, 1},
		{0, 1, 0, 1, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 1, 1, 0},
	}
	fmt.Printf("Output mineweeper %d\n", mineSweeperDetector(m1))
	fmt.Printf("Output mineweeper %d\n", mineSweeperDetector(m2))
}
