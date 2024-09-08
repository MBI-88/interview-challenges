package main

import (
	"fmt"
	"strconv"
	"strings"
)

func intToArrStr(t int) []string {
	return strings.Split(strconv.Itoa(t), "")
}

func checkPalindrome(t int) bool {
	arr := intToArrStr(t)

	for i,j := 0,len(arr) - 1; i < j ; i,j = i+1,j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}

func largestPalindromeNumber() int {
	var (
		maximun   = 101101 // smallest
		backtrack func(a,b int)
	)

	backtrack = func(a, b int) {
		if a*b > maximun && checkPalindrome(a*b) {
			maximun = a*b
		}

		if a > 100 && b > 100 {
			backtrack(a,b-1)
			backtrack(a-1,b)
		}
	}

	backtrack(999, 999)
	return maximun
}

func main() {
	fmt.Printf("Largest palindrome %d\n", largestPalindromeNumber())
}