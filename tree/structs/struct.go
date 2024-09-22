package structs

import (
	"fmt"
	"math"
)

// Tree struct
type Tree struct {
	data        int
	left, right *Tree
}

// Append function. Add a node to the Tree.
func Append(t *Tree, data int) *Tree {
	if t == nil {
		t := new(Tree)
		t.data = data
		return t
	}
	if data < t.data {
		t.left = Append(t.left, data)
	} else {
		t.right = Append(t.right, data)
	}
	return t
}

// IsEmpty function.
func IsEmpty(t *Tree) bool {
	if t == nil {
		return true
	}
	return false
}

// DelNode function. Delete a node
func DelNode(t *Tree, data int) {
	if t == nil {
		return
	}
	if t.data == data {
		t = nil
		return
	}
	DelNode(t.left, data)
	DelNode(t.right, data)
}

// Depth function. Return the tree depth
func Depth(t *Tree) int {
	if t == nil {
		return 0
	}
	return 1 + int(math.Max(float64(Depth(t.left)), float64(Depth(t.right))))
}

// PreOrder function.Run pre-order
func PreOrder(t *Tree) {
	if t == nil {
		return
	}
	fmt.Printf("%d->", t.data)
	PreOrder(t.left)
	PreOrder(t.right)
}

// PosOrder function.Run pos-order
func PosOrder(t *Tree) {
	if t == nil {
		return
	}
	PosOrder(t.left)
	PosOrder(t.right)
	fmt.Printf("%d->", t.data)
}

// InOrder function.Run in-order
func InOrder(t *Tree) {
	if t == nil {
		return
	}
	InOrder(t.left)
	fmt.Printf("%d->", t.data)
	InOrder(t.right)
}

// Leave function. Return leave's total
func Leave(t *Tree) int {
	if t == nil {
		return 0
	}
	if t.left == nil && t.right == nil {
		return 1
	}
	return Leave(t.left) + Leave(t.right)
}

// FindNode function return data fund it
func FindNode(t *Tree, data int) int {
	if t == nil {
		return 0
	}
	if t.data == data {
		return t.data
	}
	return FindNode(t.left, data) | FindNode(t.right, data)
}

// ShowTree function print Tree
func ShowTree(t *Tree, count int) {
	if t == nil {
		return
	}
	ShowTree(t.right, count+1)
	for i := 0; i < count; i++ {
		fmt.Printf("   ")
	}
	fmt.Printf("%d\n", t.data)
	ShowTree(t.left, count+1)
}
