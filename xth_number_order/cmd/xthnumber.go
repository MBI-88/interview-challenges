package cmd

import (
	"sort"
	"strconv"
	"strings"
)


func GetX(number int, arr []int) int {
	var xth int 

	sort.Ints(arr)
	xth = arr[number - 1]

	return  xth
}


func StrToInt(arr string) []int {
	var arrInts []int 

	str := strings.Split(arr, ",")

	for _, s := range str {
		n, err := strconv.Atoi(s)
		if err == nil {
			arrInts = append(arrInts, n)
		}
	}
	return arrInts
}

