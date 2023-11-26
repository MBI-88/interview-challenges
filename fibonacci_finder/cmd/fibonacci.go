package cmd

import "strconv"


func FiboFinder(k int) int {
	if k < 0 {
		return 0
	}
	var (
		arrfibo []int
		count int 
		n1 int 
		n2 int = 1
	) 

	for count < k {
		nth := n1 + n2 
		arrfibo = append(arrfibo, n1)
		n1 = n2 
		n2 = nth
		count++
	}

	return arrfibo[k - 1]
}

// Converts from str to int
func StrToInt(v string) int {
	number, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return number
}