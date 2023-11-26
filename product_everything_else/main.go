package main

import (
	"fmt"
	"os"
	"product/cmd"
	"strconv"
)

func helpInput(data []string) {
	var numberArr []int 
	for _, n := range data {
		number, _ := strconv.Atoi(n)
		numberArr = append(numberArr, number)
	}

	result := cmd.ProductOfOthers(numberArr)
	fmt.Printf("Result => %v\n", result)

}

func main() {
	arr := os.Args[1:]
	fmt.Println(arr)
	helpInput(arr)
}