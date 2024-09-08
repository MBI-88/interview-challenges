package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	r *int
	k *int
)

func showCombinatory(com [][]int) {
	for i := 0; i < len(com); i++ {
		fmt.Printf("%d\n", com[i])
	}
}

func combinatoryR(arr int, k int) {
	var (
		com       [][]int
		track     []int
		backtrack func(index int)
	)

	backtrack = func(index int) {
		if len(track) == k {
			temp := make([]int, k)
			copy(temp, track)
			com = append(com, temp)
			return
		}

		for i := index; i <= arr; i++ {
			track = append(track, i)
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}

	backtrack(1)
	showCombinatory(com)
}

func combinatoryI(arr int, k int) {
	var (
		com   [][]int
		track []int
		count = arr
		a     int
	)

	for i := 1; i <= arr; i++ {
		track = append(track, i)
		a = i
		for  count >= 0 {
			for j := a + 1; j <= arr; j++ {
				track = append(track, j)
				if len(track) == k {
					t := make([]int, k)
					copy(t, track)
					com = append(com, t)
					track = track[:len(track)-1]
				}
			}

			count--
			a++
			track = []int{i}
		}
		count = arr
		track = []int{}
		

	}

	showCombinatory(com)

}

func init() {
	r = flag.Int("range", 0, "number space")
	k = flag.Int("k", 2, "posible combinations")
}

func main() {
	flag.Parse()

	start := time.Now()
	//combinatoryR(*r, *k)
	//fmt.Printf("Time taken %d ns\n", time.Since(start).Nanoseconds())
	
	//start = time.Now()
	combinatoryI(*r, *k)
	fmt.Printf("Time taken %d ns\n", time.Since(start).Nanoseconds())

}
