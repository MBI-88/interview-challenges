package cmd

import (
	"fmt"
	"strconv"
)

type linkedList struct {
	data int
	next *linkedList
}

func (l *linkedList) Push(data int) *linkedList {
	if l == nil {
		l = new(linkedList)
		l.data = data
		return l
	}
	l.next = l.next.Push(data)

	return l
}

func (l *linkedList) ReverseLinkedList() {
	if l == nil {
		return
	}
	l.next.ReverseLinkedList()
	fmt.Printf("%d -> ", l.data)

}

func (l *linkedList) PrintLinkedList() {
	if l == nil {
		return
	}
	fmt.Printf("%d -> ", l.data)
	l.next.PrintLinkedList()
}

func (l linkedList) StrToInt(strArray []string) []int {
	var intArray []int
	for _, str := range strArray {
		number, er := strconv.Atoi(str)
		if er == nil {
			intArray = append(intArray, number)
		}

	}

	return intArray

}
