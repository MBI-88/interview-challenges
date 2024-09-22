package structs

import (
	"fmt"
)


// Queue struct
type Queue struct {
	data int
	next *Queue
}

// Push function return Queue pointer
func Push(q *Queue, data int) *Queue {
	if q == nil {
		n := new(Queue)
		n.data = data
		return n
	 }
	 q.next = Push(q.next,data)
	return q
}

// Pop function return Queue pointer
func Pop(q *Queue) *Queue {
	if q == nil {return q }
	n := new(Queue)
	n = q.next 
	q = nil 
	return n
}


// Find function return int
func Find(q *Queue, data int) int {
	if q == nil {return 0}
	if q.data == data {
		return q.data
	}
	return Find(q.next,data)
}

// Print function print Queue
func Print(q *Queue) {
	if q == nil {return }
	fmt.Printf("%d->",q.data)
	Print(q.next)
}

// Len function return Queue size
func Len(q *Queue) int {
	if q == nil {return 0}
	return 1 + Len(q.next)
}
