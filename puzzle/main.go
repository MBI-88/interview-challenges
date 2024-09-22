package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
)

/*

dashboard
[
	[ F, O, O, F]
	[ O, O, F, O]
	[ D, O, O, F]
	[ O, O, O, H]
]

F = gift
D = dog
H = Home
O = step

*/

var (
	steps = [4][2]int{
		{1, 0}, // y,x
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	att *int
)

func ramdonStep() []int {
	p := rand.Intn(4)
	return steps[p][:]
}

// next step. If the next step is into the matriz range and it not used yet
func nextStep(y, x, matrizlen int, count *int, coords [][]int) []int {
	selected := ramdonStep()
	nY := y + selected[0]
	nX := x + selected[1]

	// Validate 1
	if (nY >= 0 && nY < matrizlen) && (nX >= 0 && nX < matrizlen) {
		// Validate 2
		for _, pos := range coords {
			if pos[0] == nY && pos[1] == nX {
				return []int{y, x}
			}
		}
		*count++
		return []int{nY, nX}
	} else {
		return []int{y, x}
	}

}

func checkRewards(rewards *int, trewards, y, x int, dash [][]rune) []int {
	if dash[y][x] == 'F' {
		*rewards++
		return []int{y, x}
	} else if dash[y][x] == 'H' && *rewards == trewards {
		*rewards++
	} else if dash[y][x] == 'H' && *rewards != trewards {
		*rewards = -1
	}
	return []int{-1, -1}
}

func findRewards(dash [][]rune) int {
	var t int

	for y := 0; y < len(dash); y++ {
		for x := 0; x < len(dash[y]); x++ {
			if dash[y][x] == 'F' {
				t++
			}
		}
	}

	return t
}

func findDog(dash [][]rune) []int {
	var dog []int

	for y := 0; y < len(dash); y++ {
		for x := 0; x < len(dash[y]); x++ {
			if dash[y][x] == 'D' {
				dog = append(dog, y, x)
			}
		}
	}

	return dog
}

func checkDashboard(dash [][]rune) int {
	var size int
	ylen := len(dash)

	for y := 0; y < ylen; y++ {
		for x := 0; x < len(dash[y]); x++ {
			size++
		}
		if size != ylen {
			return 0
		}
		size = 0
	}

	return ylen
}

func bestTrajectory(track []int) {
	var min = math.MaxInt
	for _, p := range track {
		if p < min {
			min = p
		}
	}
	fmt.Printf("The best trajectory with total steps %d\n", min)
}

func play(dashboard [][]rune, attempts int) {
	var (
		cellVisited  [][]int    // add coords
		rewards      = new(int) // rewards must be 5 including 'H' , home
		trajectories []int      // trajectory completed only if all rewards gotten
		count        = new(int)
		errorCount   int
	)
	totalRewards := findRewards(dashboard)
	matrizlen := checkDashboard(dashboard)
	if matrizlen == 0 {
		fmt.Println("Matriz must be square")
		return
	}
	fmt.Println("Playing...")

	// trajectory
start:
	dog := findDog(dashboard)
	cellVisited = append(cellVisited, []int{dog[0], dog[1]})

	for attempts > 0 {
		// loop
		currentY := dog[0]
		currentX := dog[1]
		dog = nextStep(currentY, currentX, matrizlen, count, cellVisited)
		if dog[0] != currentY || dog[1] != currentX {
			cell := checkRewards(rewards, totalRewards, dog[0], dog[1], dashboard)
			if cell[0] != -1 && cell[1] != -1 {
				cellVisited = append(cellVisited, cell)
			}
			// validation
			if *rewards == totalRewards + 1 { // Successful trajectory
				trajectories = append(trajectories, *count)
				cellVisited = [][]int{}
				*rewards = 0
				*count = 0
				goto start

			} else if *rewards == -1 { // Error trajectory
				*count = 0
				*rewards = 0
				cellVisited = [][]int{}
				goto start

			}

			attempts--
		} else if errorCount > attempts * 10 {
			*count = 0
			*rewards = 0
			cellVisited = [][]int{}
			goto start
		}

		errorCount++
	}

	fmt.Println(trajectories)
	bestTrajectory(trajectories)
}

func init() {
	att = flag.Int("attempts", 0, "total attempts")
}

func main() {
	flag.Parse()
	dash1 := [][]rune{
		{'F', 'O', 'F'},
		{'O', 'F', 'F'},
		{'D', 'O', 'H'},
	}
	play(dash1, *att)
}
