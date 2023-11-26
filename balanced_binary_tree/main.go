package main


import (
	"balanced_tree/cmd"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Coverts from array string to array number
func StrToInts(str []string) []int {
	var arrInst []int

	for _, item := range str {
		number, err := strconv.Atoi(item)
		if err == nil {
			arrInst = append(arrInst, number)
		}
	}

	return arrInst
}

// Feeds the AVL
func BalancedBST(orderedList []int) int {
	avl := cmd.NewBalancedTree(orderedList[0])
	for _, item := range orderedList[1:] {
		avl = avl.Add(item)
	}
	avl.PrintAVL(1)
	return avl.RootNode()
}


func operation() {
	input := os.Args[1:]
	arrayInts := StrToInts(input)
	sort.Ints(arrayInts)
	fmt.Printf("Root node %d\n", BalancedBST(arrayInts))

}

func main() {
	operation()
}



