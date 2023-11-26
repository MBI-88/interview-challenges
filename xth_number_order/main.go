package main

import (
	"flag"
	"fmt"
	"os"
	"xth-number/cmd"
)

var (
	number *int
	array  *string
)

func init() {
	number = flag.Int("n", 0, "set the x number")
	array = flag.String("ar", "", "set the array of number separed by (,)")
	flag.Usage = func() {
		info := fmt.Sprint("[*] Find the Xth Number in Order\n")
		info += "\tvariable -n is the position for finding the xth number\n"
		info += "\tvariable -ar is the array of number\n"
		fmt.Fprintf(os.Stderr, "%s", info)
		flag.PrintDefaults()
	}

}

func opertation() {
	flag.Parse()
	arrNumber := cmd.StrToInt(*array)
	result := cmd.GetX(*number, arrNumber)
	fmt.Printf("The integer corresponding to Xth in the array is %d\n", result)
}

func main() {
	opertation()

}
