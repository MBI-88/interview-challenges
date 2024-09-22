package main

import (
	"fmt"
)

func findOriginPos(arr []any, item any, pos *[]int) int {
	var fp = false
	for i, el := range arr {
		if el == item {
			for _, x := range *pos {
				if x == i {
					fp = true
					break
				}
			}
			if !fp {
				*pos = append(*pos, i)
				return i
			}
			fp = false
		}
	}

	return -1
}

func getItemsAt(arr []any, flag string) []any {
	var (
		result   []any
		rmap     = make([]any, len(arr))
		reversed []any
		pos      = new([]int)
	)

	for index := len(arr) - 1; index >= 0; index-- {
		reversed = append(reversed, arr[index])
	}

	for index := 1; index <= len(arr); index++ {
		if flag == "odd" && index%2 != 0 {
			if rs := findOriginPos(arr, reversed[index-1], pos); rs != -1 {
				rmap[rs] = reversed[index-1]
			}
		} else if flag == "even" && index%2 == 0 {
			if rs := findOriginPos(arr, reversed[index-1], pos); rs != -1 {
				rmap[rs] = reversed[index-1]
			}
		}
	}

	for _, item := range rmap {
		if item != nil {
			result = append(result, item)
		}
	}

	return result
}

func main() {

	fmt.Println(getItemsAt([]any{2, 4, 6, 8, 10}, "odd")...) // -> [2, 6, 10]
	fmt.Println(getItemsAt([]any{"E", "D", "A", "B", "I", "T"}, "even")...) // -> ["E", "A", "I"]
	fmt.Println(getItemsAt([]any{")", "(", "*", "&", "^", "%", "$", "#", "@", "!"}, "even")...) // -> [")", "*", ^", "$", "@"]
	fmt.Println(getItemsAt([]any{"A", "R", "B", "I", "T", "R", "A", "R", "I", "L", "Y"}, "even")...) //-> ["R", "I", "R", "R", "L"]
}
