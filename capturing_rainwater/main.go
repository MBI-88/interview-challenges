package main

import (
	"fmt"
	"os"
	"rainwater/cmd"
)

func opetation() {
	input := os.Args[1:]
	inputIn := cmd.StringToInt(input)
	result := cmd.CapturingRainwater(inputIn)
	fmt.Printf("Rainwater captured %d\n", result)
}

func main() {
	opetation()
}
