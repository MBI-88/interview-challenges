package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)


func kapsack(wei, val []int, cap int) int {
	var (
		maximun   int
		backtrack func(i, tw, tv int)
	)

	if len(val) != len(wei) {
		panic(fmt.Errorf("Error with array length"))
	}

	backtrack = func(index, tw, tv int) {
		if tw <= cap && tv > maximun {
			maximun = tv
		}
		
		if tw > cap {
			return
		}

		for i := index; i < len(val); i++ {
			backtrack(i + 1, tw + wei[i], tv + val[i])
		}
	}

	backtrack(0, 0, 0)
	return maximun
}



func strToInt(w, v string) ([]int, []int) {
	var (
		weights []int
		values      []int
	)

	strW := strings.Split(w, ",")
	strV := strings.Split(v, ",")

	if len(strV) != len(strW) {return nil,nil}

	for i := 0; i < len(strV); i++ {
		if tempW,err := strconv.Atoi(strW[i]); err == nil {
			weights = append(weights, tempW)
		}

		if tempV, err:= strconv.Atoi(strV[i]); err == nil {
			values = append(values, tempV)
		}
	}
	return weights, values
}

var (
	weights *string 
	values *string
	unit *int
)


func init(){
	weights = flag.String("weights", "", "set weights")
	values = flag.String("values", "", "set values respect to weights")
	unit = flag.Int("unit", 10, "backpack cantity")
}

func main() {
	flag.Parse() 
	w, v := strToInt(*weights, *values)
	fmt.Printf("Kapsack result %v\n", kapsack(w, v, *unit))
}