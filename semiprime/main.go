package main

import (
	"fmt"
	"os"
	"semiprime/cmd"
	"strconv"
)


func main() {
	val := os.Args[1]
	number, _ := strconv.Atoi(val)
	fmt.Printf("Semiprime numbers count %d\n", cmd.SemiprimeCount(number))
}