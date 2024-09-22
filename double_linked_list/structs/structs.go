package structs

import (
	"fmt"
)

// DoubleLinkedList struct
type DoubleLinkedList struct {
	data       int
	prev, next *DoubleLinkedList
}

// Append function add an element in the list
func Append(d *DoubleLinkedList, data int) *DoubleLinkedList {
	if d == nil {
		n := new(DoubleLinkedList)
		n.data = data
		return n
	}
	d.next = Append(d.next, data)
	d.next.prev = d
	return d
}

// Delete function remove an element from the list
func Delete(d *DoubleLinkedList, data int) *DoubleLinkedList {
	if d == nil {
		return d
	}
	if d.data == data && d.prev == nil && d.next != nil {
		d.next.prev = nil
		return d.next
	}else if d.data == data && d.prev != nil && d.next != nil {
		d.next.prev = d.prev
		return d.next
	}else if d.data == data && d.next == nil && d.prev != nil {
		return d.next
	}else if d.data == data && d.prev == nil && d.next == nil {
		return nil
	}

	d.next = Delete(d.next, data)
	return d
}

// Find function find an element in the list
func Find(d *DoubleLinkedList, data int) int {
	if d == nil {return 0}
	if d.data == data {
		return d.data
	}
	return Find(d.next,data)
}

// Len function return list size
func Len(d *DoubleLinkedList) int {
	if d == nil {return 0}
	return 1 + Len(d.next)
}

// Print function print the list
func Print(d *DoubleLinkedList) {
	if d == nil {return }
	fmt.Printf("<-%d->",d.data)
	Print(d.next)
}

// Clean function clean the list
func Clean(d *DoubleLinkedList) *DoubleLinkedList {
	if d == nil {return d}
	d = nil
	return d
}
