package main

import (
	"fibonacci/cmd"
	"fmt"
	"os"
)


func opertion() {
	input := os.Args[1]
	number := cmd.StrToInt(input)
	result := cmd.FiboFinder(number)
	fmt.Printf("Fibonacci nth %d\n", result)
}

func main() {
	opertion()

}