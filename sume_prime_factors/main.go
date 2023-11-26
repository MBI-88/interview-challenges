package main

import (
	"fmt"
	"os"
	"prime-factors/cmd"
	"strconv"
)

func operation() {
	input, err := strconv.Atoi(os.Args[len(os.Args) - 1])
	if err != nil {
		panic(err)
	}
	result := cmd.SumePrimeFactors(input)
	fmt.Printf("Sum of prime factors of %d => %d\n", input, result)

}

func main() {
	operation()
}
