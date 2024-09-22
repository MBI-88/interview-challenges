package structs

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

// CircularList struct
type CircularList struct {
	head *node
	tail *node
	size int
}

func (c *CircularList) isEmpty() bool {
	return c.size == 0
}

// Append method add an element in the list
func (c *CircularList) Append(data int) {
	n := &node{data, nil}
	if c.isEmpty() {
		c.head = n
		n.next = n
		c.tail = n
	} else {
		c.tail.next = n
		c.tail = n
		c.tail.next = c.head
	}
	c.size++
}

// Delete function remove an element from the list
func (c *CircularList) Delete(data int) {
	if c.isEmpty() {
		return
	}
	if c.head.data == data {
		c.head = c.head.next
		c.tail.next = c.head

	} else if c.tail.data == data {
		newtail := c.head
		for newtail != c.tail {
			newtail = newtail.next
		}
		newtail.next = c.head
		c.tail = newtail

	} else {
		current := c.head
		for current.data != data {
			current = current.next
		}
		current.data = current.next.data
		current.next = current.next.next

	}

	c.size--

}

// Print method print the list
func (c *CircularList) Print() {
	if c.isEmpty() {
		return
	}
	current := c.head
	for i := 0; i < c.size; i++ {
		fmt.Printf("%d->", current.data)
		current = current.next
	}

}

// Len method return size list
func (c *CircularList) Len() int {
	return c.size
}

// Find method find an element int the list
func (c *CircularList) Find(data int) int {
	if c.isEmpty() {
		return 0
	}
	found := c.head
	for found.data != data {
		found = found.next
	}
	return found.data
}

// Clean method remove all elements in the list
func (c *CircularList) Clean() {
	if c.isEmpty() {
		return
	}
	c.size = 0
}
