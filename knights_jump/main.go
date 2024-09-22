package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var (
	chessBoard = [8][8]int{
	   	{1, 2, 3, 4, 5, 6, 7, 8},         // 0
		{9, 10, 11, 12, 13, 14, 15, 16},  // 1
		{17, 18, 19, 20, 21, 22, 23, 24}, // 2
		{25, 26, 27, 28, 29, 30, 31, 32}, // 3
		{33, 34, 35, 36, 37, 38, 39, 40}, // 4
		{41, 42, 43, 44, 45, 46, 47, 48}, // 5
		{49, 50, 51, 52, 53, 54, 55, 56}, // 6
		{57, 58, 59, 60, 61, 62, 63, 64}, // 7
		//A,  B,  C,  D,  E,  F,  G,  H,
	}
	chessIntToString = make(map[int]string)
	chessStringoInt  = make(map[string]int)
	letters          = [8]string{"A", "B", "C", "D", "E", "F", "G", "H"}
	movements        = [8][2]int{
		{2, -1},  // [x,y] always start by y
		{1, -2},  // [x,y] always start by y
		{-1, -2}, // [x,y] always start by y
		{-2, -1}, // [x,y] always start by y
		{2, 1},   // [x,y] always start by y
		{1, 2},   // [x,y] always start by y
		{-1, 2},  // [x,y] always start by y
		{-2, 1},  // [x,y] always start by y
	}
	
)

func fillMaps() {
	for step := range 8 {
		chessIntToString[step] = letters[step]
		chessStringoInt[letters[step]] = step
	}
}

func getCoords(pos string) (int, int) {
	fillMaps()
	coords := strings.Split(pos, "")
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		panic(err)
	}
	x := chessStringoInt[coords[0]]
	return x, y
}

func findCoord(pos int) string {
	var (
		stop   = false
		result string
	)
	for x := 1; x < len(chessBoard); x++ {
		for y := 1; y < len(chessBoard); y++ {
			if chessBoard[x][y] == pos {
				strX := chessIntToString[x]
				result = fmt.Sprintf("%s%d", strX, y)
				stop = true
				break
			}
		}
		if stop {
			break
		}
	}
	return result
}

func knightsJump(pos string) string {
	var trackPos []string
	x, y := getCoords(pos)
	for _, step := range movements {
		if (x+step[0] > 0 && x+step[0] <= 7) && (y+step[1] > 0 && y+step[1] <= 7) {
			post := chessBoard[x+step[0]][y+step[1]]
			trackPos = append(trackPos, findCoord(post))
		}
	}
	sort.Strings(trackPos)
	return strings.Join(trackPos, ",")
}

func main() {
	fmt.Println(knightsJump("F4"))
	fmt.Println(knightsJump("E2"))
	fmt.Println(knightsJump("A1"))
}
