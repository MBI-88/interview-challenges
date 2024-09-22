package main

import (
	"fmt"
	"strings"
)



func freqCount(str, obj string) map[int]int {
	var (
		freq  = make(map[int]int)
		array = strings.Split(str, "")
		level = -1
	)

	for _, item := range array {
		if item == "[" {
			level++
			freq[level] += 0
		}else if item == "]" {
			level--
		}

		if item == obj {
			freq[level]++
		}
	}

	return freq
}

func main() {
	fmt.Println(freqCount("[1, [2], 1, [[2]], 1, [[[2]]], 1, [[[[2]]]]]", "2")) 
	// freqCount("[1,4,4,[1,1,[1,2,1,1]]]", "1")
	//freqCount("[1, 5, 5, [5, [1, 2, 1, 1], 5, 5], 5, [5]]", "5")
	//freqCount("[1, [2], 1, [[2]], 1, [[[2]]], 1, [[[[2]]]]]", "2")

}
