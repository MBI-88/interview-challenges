package main

import (
	"fmt"
	"math"
)

var (
	path1 = [4][4]string{
		{"F", "", "", ""},
		{"", "D", "", "F"},
		{"", "F", "", "F"},
		{"", "", "", "H"},
	}
	path2 = [4][4]string{
		{"", "F", "", ""},
		{"F", "", "D", ""},
		{"", "", "", "F"},
		{"", "F", "", "H"},
	}
	path3 = [4][4]string{
		{"", "F", "", "F"},
		{"D", "", "F", ""},
		{"", "", "", ""},
		{"F", "", "", "H"},
	}
)

func getPositions(m [4][4]string) ([][]int, [][]int, [][]int) {
	var (
		rewards [][]int
		dog     [][]int
		home    [][]int
	)

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			if m[i][j] == "F" {
				rewards = append(rewards, []int{i, j})
			} else if m[i][j] == "H" {
				home = append(home, []int{i, j})
			} else if m[i][j] == "D" {
				dog = append(dog, []int{i, j})
			}
		}
	}
	return rewards, dog, home
}

func makeSum(actual, next []int) int {
	xSum := math.Abs(float64(next[0] - actual[0]))
	ySum := math.Abs(float64(next[1] - actual[1]))
	return int(xSum + ySum)
}

func bestSteps(m [4][4]string) int {
	var (
		rewards, dog, home = getPositions(m)
		steps              int
		distances       [][][]int
		combinations    [][][]int
		backtrack          func(actual int)
		acc                []int
		minimun            = math.MaxInt64
	)

	backtrack = func(actual int) {
		if actual == len(rewards) {
			temp := make([][]int, len(rewards))
			copy(temp, rewards)
			distances = append(distances, temp)
			return
		}

		for i := actual; i < len(rewards); i++ {
			backtrack(actual + 1)
			rewards[actual], rewards[i] = rewards[i], rewards[actual]
		}
	}

	backtrack(0)

	for i := 0; i < len(distances); i++ {
		temp := [][]int{}
		temp = append(temp, dog...)
		temp = append(temp, distances[i]...)
		temp = append(temp, home...)
		combinations = append(combinations, temp)

	}

	for i := 0; i < len(combinations); i++ {
		for x := 0; x < len(combinations[i])-1; x++ {
			steps += makeSum(combinations[i][x], combinations[i][x+1])
		}
		acc = append(acc, steps)
		steps = 0
	}
	for _, ac := range acc {
		if ac < minimun {
			minimun = ac
		}
	}

	return minimun
}

func main() {
	fmt.Printf("Best path1 %d\n", bestSteps(path1))
	fmt.Printf("Best path2 %d\n", bestSteps(path2))
	fmt.Printf("Best path3 %d\n", bestSteps(path3))
}
