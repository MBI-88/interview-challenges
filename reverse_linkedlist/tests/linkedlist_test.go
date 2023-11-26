package tests

import (
	"fmt"
	"io"
	"linkedlist/cmd"
	"os"
	"testing"
)

func TestLinkedList(t *testing.T) {
	source := []int{1,2,3,4,5}
	r, w, _ :=  os.Pipe()
	stdout := os.Stdout
	os.Stdout = w 

	list := cmd.NewLinkedLis()
	for _, item := range source {
		list.Push(item)
	}

	list.ReverseLinkedList()
	w.Close()

	ouput, _ := io.ReadAll(r)
	stdOuput := string(ouput)
	os.Stdout = stdout 

	if stdOuput != fmt.Sprint("5 -> 4 -> 3 -> 2 -> 1 -> 0 -> ") {
		t.Fatalf("Error in stdout %s",stdOuput)
	}
}