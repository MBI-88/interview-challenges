package structs

import (
	"fmt"
)

// PriotityQueue struct
type PriotityQueue struct {
	data int
	next *PriotityQueue
}

// Push function return QueuePriotity pointer
func Push(q *PriotityQueue, data int) *PriotityQueue {
	if q == nil {
		n := new(PriotityQueue)
		n.data = data
		return n

	} else if data > q.data {
		temp := q.data
		q.data = data
		q.next = Push(q.next,temp)
		return q

	} 
	q.next = Push(q.next, data)
	return q
	
}

// Pop function return the first node
func Pop(q *PriotityQueue) *PriotityQueue {
	if q == nil {return q}
	n := q.next 
	q = nil
	return n
}

// Find function return a data found
func Find(q *PriotityQueue,data int) int {
	if q == nil {return 0}
	if q.data == data {
		return q.data
	}
	return Find(q.next,data)
}

// Print function print Queue
func Print(q *PriotityQueue) {
	if q == nil {return }
	fmt.Printf("%d->",q.data)
	Print(q.next)
}

// Len function return queue size
func Len(q *PriotityQueue) int {
	if q == nil {return 0}
	return 1 + Len(q.next)
}