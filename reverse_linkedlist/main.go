package main

import (
	"fmt"
	"linkedlist/cmd"
	"os"
)




func operation() {
	arr := os.Args[1:]
	list := cmd.NewLinkedLis()
	arrInts := list.StrToInt(arr)
	for _, item := range arrInts {
		list.Push(item)
	}

	fmt.Println("Linked list ordered")
	list.PrintLinkedList()
	fmt.Println()
	fmt.Println("Linked list reversed")
	list.ReverseLinkedList()
	fmt.Println()
	
}

func main() {
	operation()

}