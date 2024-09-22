package structs

import (
	"fmt"
)

// List struct 
type List struct {
	data int 
	next *List
}

// Append function add an element in the list
func Append(l *List,data int) *List {
	if l == nil {
		n := new(List)
		n.data = data
		return n
	}
	l.next = Append(l.next,data)
	return l
}

// Delete function remove an element from the list
func Delete(l *List, data int) *List {
	if l == nil {return l}
	
	if l.data == data {
		return l.next
	}
	l.next = Delete(l.next,data)
	return l
}

// Find function return an element found
func Find(l *List,data int) int {
	if l == nil {
		return 0
	}else if l.data == data {
		return l.data
	}
	return Find(l.next,data)
}

// Print function print the list
func Print(l *List) {
	if l == nil {return }
	fmt.Printf("%d->",l.data)
	Print(l.next)
}

// Len function return list size
func Len(l *List) int {
	if l == nil {return 0}
	count := 1
	for l.next != nil {
		count++
		l = l.next
	}
	return count
}

// Clean fucntion to clean all elements in the list
func Clean(l *List) *List {
	if l == nil {return l}
	l = nil
	return l
}