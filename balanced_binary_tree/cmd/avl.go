package cmd

import (
	"fmt"
)

// AVL is a ainary search tree
type AVL struct {
	node   int
	height int
	left   *AVL
	right  *AVL
}

// Performs a right rotation on the node
func (a *AVL) rightRotation() *AVL {
	head := a.left  
	t := head.right
	head.right = a
	a.left = t
	a.height = a.max(a.left.getHeight(), a.right.getHeight()) + 1
	head.height = head.max(head.left.getHeight(), head.right.getHeight()) + 1
	return head
}

// Performs a left rotation on the node
func (a *AVL) leftRotation() *AVL {
	head := a.right
	t := head.left
	head.left = a
	a.right = t
	a.height = a.max(a.left.getHeight(), a.right.getHeight()) + 1
	head.height = head.max(head.left.getHeight(), head.right.getHeight()) + 1
	return head
}

// Prints the AVL
func (a *AVL) PrintAVL(space int) {
	if a == nil {
		return
	}

	a.left.PrintAVL(space + 1)
	for i := 0; i < space; i++ {
		fmt.Printf("  ")
	}
	fmt.Println(a.getData())
	a.right.PrintAVL(space + 1)

	fmt.Println()

}

// Returns maximun of two integers
func (a AVL) max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Inserts a data itno the AVL
func (a *AVL) Add(data int) *AVL {
	if a == nil {
		a := &AVL{data, 1, nil, nil}
		return a
	}

	if data < a.getData() {
		a.left = a.left.Add(data)
	} else if data > a.getData() {
		a.right = a.right.Add(data)
	}

	
	a.height = 1 + a.max(a.left.getHeight(), a.right.getHeight())
	fe := a.Bf()

	if fe > 1 {
		if data < a.left.getData() {
			return a.rightRotation()
		} else if data > a.left.getData() {
			a.left = a.left.leftRotation()
			return a.rightRotation()
		}
	}

	if fe < -1 {
		if data > a.right.getData() {
			return a.leftRotation()
		} else if data < a.right.getData() {
			a.right = a.right.rightRotation()
			return a.leftRotation()
		}
	}
	
	return a
}


// Returns tree's height
func (a *AVL) getHeight() int {
	if a == nil {
		return 0
	}
	return a.height
}

// Returns tree's data
func (a *AVL) getData() int {
	if a == nil {
		return 0
	}
	return a.node
}

// Returns balance factor
func (a *AVL) Bf() int {
	if a == nil {
		return 0
	}
	return a.left.getHeight() - a.right.getHeight()
}

// Returns root node
func (a AVL) RootNode() int {
	return a.node
}