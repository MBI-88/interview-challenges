package main

// ["afdfdfffef","jeffaa"]

import (
	"fmt"
	"os"
)

func minWindows(Array []string) (result string) {
	array, pattern := Array[0], Array[1]
	freq := make(map[byte]int)
	count := 0
	minLen := len(array) * 2
	right,left := 0,0
	for _, item := range pattern {
		freq[byte(item)]++
	}

	for  right < len(array) {
		if _, ok := freq[byte(array[right])]; ok {
			freq[byte(array[right])]--
			if freq[byte(array[right])] >= 0 {
				count++
			}
		}
		for  count == len(pattern) {
			if _, ok := freq[byte(array[left])]; ok {
				if right-left+1 < minLen {
					minLen = right - left + 1
					result = array[left : right+1]
				}
				freq[byte(array[left])]++
				if freq[byte(array[left])] > 0 {
					count--
				}
			}
			left++

		}
		right++

	}
	return result
}

func main() {
	args := os.Args[1:]
	result := minWindows(args)
	fmt.Println(result)

}
