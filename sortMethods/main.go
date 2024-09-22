package main

import (
	"flag"
	"fmt"
	"methods/bubble"
	"methods/merge"
	"methods/quick"
	"methods/insertion"
	"os"
	"strconv"
	"strings"
)

var md *string

func init() {
	md = flag.String("md", "bubble", "select method to sort")
	flag.Usage = func() {
		info := "[*] Flags: -md (select method)\nExameple:-md=bubble\ndefault bubble"
		info += "[*] Methods: bubblesort,quicksort,mergesort\n"
		fmt.Fprintf(os.Stderr, "%s", info)
		flag.PrintDefaults()
	}
}

func cleanInput(s string) (input []int) {
	array := strings.Split(s, ",")
	for _, i := range array {
		val, er := strconv.Atoi(i)
		if er != nil {
			panic("Error in data convertion")
		}
		input = append(input, val)
	}
	return input
}

func userInterface() {
	arg := flag.Arg(0)
	if len(arg) == 0 {return}
	input := cleanInput(arg)
	switch *md {
	case "quick":
		quick := quick.Quick{Data: input}
		quick.Sort()
		quick.Print()
		fmt.Println()
	case "merge":
		merge := merge.Merge{Data: input}
		merge.Sort()
		merge.Print()
		fmt.Println()
	case "insertion":
		insertion := insertion.Insertion{Data:input}
		insertion.Sort()
		insertion.Print()
		fmt.Println()
	default:
		bubble := bubble.Bubble{Data: input}
		bubble.Sort()
		bubble.Print()
		fmt.Println()

	}

}

func main() {
	flag.Parse()
	userInterface()

}
