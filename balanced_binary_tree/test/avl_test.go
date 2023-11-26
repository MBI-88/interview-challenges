package test

import (
	"balanced_tree/cmd"
	"testing"
)



func TestBST(t *testing.T) {
	payload := struct {
		array []int
		root  int
	}{
		[]int{1, 2, 3, 4, 5, 6, 7, 8}, 4,
	}
	avl := cmd.NewBalancedTree(payload.array[0])
	for _, item := range payload.array[1:] {
		avl =  avl.Add(item)
	}

	if payload.root != avl.RootNode() {
		t.Fatalf("Expected %d recived %d", payload.root, avl.RootNode())
	}
}
