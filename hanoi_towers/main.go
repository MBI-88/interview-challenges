package main

/*
Hanoi towers solution
Rules:
-> only one disc by acction
-> discs with greater weight never on top of smaller discs

*/

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

// Recursive
func hanoiTowerRecurtion(disc int, towerA, towerB, towerC string) int {
	if disc == 0 {
		return 1
	}

	hanoiTowerRecurtion(disc-1, towerA, towerC, towerB)
	fmt.Printf("dic %d from tower %s to tower %s\n", disc, towerA, towerC)
	hanoiTowerRecurtion(disc-1, towerB, towerA, towerC)
	return disc
}

// Iterative 4 disc supported (It's posible to test every 4 jumps )

func hanoiTowerIteration(disc int) {
	if disc == 0 {
		return
	}

	towerA := "A"
	towerB := "B"
	towerC := "C"
	c := -1
	countUp := 0
	movesTotal := int(math.Pow(2, float64(disc))) - 1

	for movesTotal > 0 {
		if c == -1 {
			for i := disc - 1; i >= 1; i-- {
				towerB, towerC = towerC, towerB
				c = i

			}
		}
	op1:
		if c-1 > 0 {
			towerB, towerC = towerC, towerB
		}

		switch countUp {
		case 3:
			fmt.Printf("operation from tower %s to tower %s\n", towerC, towerA)
		case 11:
			fmt.Printf("operation from tower %s to tower %s\n", towerC, towerA)
		case 15:
			fmt.Printf("operation from tower %s to tower %s\n", towerC, towerA)
		case 19:
			fmt.Printf("operation from tower %s to tower %s\n", towerC, towerA)

		default:
			fmt.Printf("operation from tower %s to tower %s\n", towerA, towerC)
		}

		if c-1 > 0 {
			towerA, towerB = towerB, towerA
			c--
			countUp++
			movesTotal--
			goto op1
		}

		c++
		countUp++
		movesTotal--
	}

}

func main() {
	flag.Parse()
	args := flag.Arg(0)
	if len(args) > 1 {
		fmt.Printf("Too many argumes, only one is accepted\n")
	} else {
		disc, err := strconv.Atoi(args)
		if err != nil {
			panic("Argument is not a number")
		}
		//hanoiTowerRecurtion(disc, "A", "B", "C")
		hanoiTowerIteration(disc)
	}
}
