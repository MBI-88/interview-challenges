package structs

import (
	"fmt"
)


// Stack struct
type Stack struct {
	data int 
	next *Stack
}

// Push function return Stack pointer
// add item to the stack
func Push(s *Stack,data int) *Stack {
	n := new(Stack)
	n.data = data 
	n.next = s 
	s = n
	return s
}

// Pop function return Stack pointer
// pop item from the stack
func Pop(s *Stack) *Stack {
	if s == nil {
		return s
	}
	n := s.next
	s = nil
	return n
}

// Len function return integer
// give Stack long
func Len(s *Stack) int {
	if s == nil {return 0}
	return 1 + Len(s.next)
}

// Find function return a number found
func Find(s *Stack, data int) int {
	if s == nil {return 0}
	if s.data == data {
		return s.data
	}
	return Find(s.next,data)
}

// Print function print the stack
func Print(s *Stack) {
	if s == nil {return }
	fmt.Printf("%d->",s.data)
	Print(s.next)
}