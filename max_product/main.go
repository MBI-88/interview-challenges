package main

import (
	"flag"
	"fmt"
	"maxproduct/cmd"
	"os"
)


var (
	k *int
)

// Inits variables 
func init() {
	k = flag.Int("k", 2, "set count attained")
	flag.Usage = func() {
		info := fmt.Sprintf("[-] Max product finder\n")
		info += "[-] Netative number in the begining not allowed\n"
		fmt.Fprintf(os.Stderr, "%s", info)
		flag.PrintDefaults()
	}
}

// Makes operation
func operation() {
	flag.Parse()
	input := flag.Args()
	arrInt := cmd.StrToInt(input)
	mxProduct := cmd.MaxProductFinder(arrInt, *k)
	fmt.Printf("Max product finder %d\n", mxProduct)

}


func main() {
	operation()
}