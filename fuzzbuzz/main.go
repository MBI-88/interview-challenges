package main

import (
	"flag"
	"fmt"
	"fuzz/cmd"
	"os"
)



var number *int

func init() {
	number = flag.Int("number", 0, "set a number")
	flag.Usage = func() {
		info := fmt.Sprintf("[*] FuzzBuzz problem")
		fmt.Fprintf(os.Stderr, "%s", info)
		flag.PrintDefaults()
	}

}

func opertion() {
	flag.Parse()
	result := cmd.FizzBuzz(*number)
	fmt.Printf("Result %v\n", result)
}

func main() {
	opertion()
}