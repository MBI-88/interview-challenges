package main

import (
	"fmt"
	"os"
	"words/cmd"
)


// Operation helper
func operation() {
	input := os.Args[1:]
	inputStr := cmd.ArrToString(input)
	result := cmd.WordReverse(inputStr)
	fmt.Printf("Input: %s\n Output: %s\n", inputStr, result)

}


func main() {
	operation()
}