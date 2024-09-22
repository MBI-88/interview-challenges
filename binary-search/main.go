package main

import "fmt"

func binarySearch(arr []int, s int) {
	if len(arr) == 1 {
		if arr[0] == s {
			fmt.Printf("Match  %d == %d\n",arr[0], s)
			return
		}
		fmt.Println("Not found")
		return
	}

	r := len(arr) - 1
	medium := int(r / 2)

	if s < arr[medium] { // go by left
		r = medium - 1
		binarySearch(arr[:r+1], s)
	}else if s > arr[medium] { // go by right
		binarySearch(arr[medium+1:], s)
	}else {
		fmt.Printf("Match  %d == %d\n",arr[medium], s)
		return
	}
}

func main() {
	var arr []int
	for item := range 100 {
		arr = append(arr, item)
	}

	binarySearch(arr, 10)
	binarySearch(arr, 5)
	binarySearch(arr, 80)
	binarySearch(arr, 101)
}