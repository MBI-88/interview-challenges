package cmd

import (
	"strconv"
)

// Finds the maximun product attained form any k integer
func MaxProductFinder(arr []int, k int) int {
	if len(arr) < k || k < 2{
		return 0
	}

	var (
		product int
		index1   int
		index2  int
		count int
	)
	post := make(map[int]bool)

	product, index1, index2 = maxBinaryProduct(arr)
	post[index1] = true
	post[index2] = true
	count = 2

	if k > 2 {
		maxNexProduct(arr, post, &product, &count, k)
	}

	return product
}

// Makes the initial binary product
func maxBinaryProduct(arr []int) (int, int, int) {
	var (
		prod  int
		index1 int
		index2 int
	)
	for i := 0; i < len(arr); i++ {
		for x := i + 1; x < len(arr); x++ {
			if arr[i]*arr[x] > prod {
				prod = arr[i] * arr[x]
				index1 = i
				index2 = x
			}
		}
	}
	return prod, index1, index2
}

// Finds the next maximun number for the maximun product
func maxNexProduct(arr []int, post map[int]bool, pd, c *int, k int) {
	for *c < k {
		found := findMaximun(arr, post)
		*pd *= found
		*c++
	}
}

// Finds maximun not contained into the postiion dictionary
func findMaximun(arr []int, dic map[int]bool ) int {
	var (
		maximun int 
		index int
	)
	for i := 0; i < len(arr); i++ {
		if !dic[i] && arr[i] > maximun {
			maximun = arr[i]
			index = i
			
		}
	}
	dic[index] = true
	return maximun
}


// Converts from str to int array
func StrToInt(arr []string) []int {
	var re []int
	for _, item := range arr {
		number, _ := strconv.Atoi(item)
		re = append(re, number)
	}
	return re
}
