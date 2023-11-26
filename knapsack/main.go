package main

import (
	"flag"
	"fmt"
	"knap/cmd"
	"os"
)

var (
	weights *string
	values *string
	limit *int
)


func init() {
	weights = flag.String("weights", "", "enter weights separated by (,)")
	values = flag.String("values", "", "enter values separated by (,)")
	limit = flag.Int("limit", 10, "enter limit, default 10") 

	flag.Usage = func() {
		info := fmt.Sprintf("[*] Knapsack problem")
		fmt.Fprintf(os.Stderr, "%s", info)
		flag.PrintDefaults()
	}
}

// Do operation
func operation(w []int, v []int, u int) {
 	reult := cmd.Knapsack(w, v, u)
	fmt.Printf("Max value recovered %d\n", reult)
}

func main() {
	flag.Parse()
	operation(cmd.StrToInt(*weights), cmd.StrToInt(*values), *limit)
	
	
}