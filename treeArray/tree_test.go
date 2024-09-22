package main


import (
	"testing"
)

func TestArrayOk1(t *testing.T){
	array := "[(3,2),(5,2),(2,4),(1,4),(6,1),(5,1)]"
	if !FindTree(array) {
		t.Error("Error TestArrayOk1")
	}
}

func TestArrayError1(t *testing.T) {
	array := "[(2,3),(1,3),(4,2),(5,2),(7,8)]"
	if FindTree(array) {
		t.Error("Error TestArrayError1")
	}
}

func TestArrayOk2(t *testing.T) {
	array := "[(6,4),(2,1),(3,1),(8,4),(4,2),(9,3)]"
	if !FindTree(array) {
		t.Error("Error TestArrayOk2")
	}
}

func TestArrayError2(t *testing.T) {
	array := "[(1,2),(3,2),(4,1),(5,3),(9,10)]"
	if FindTree(array) {
		t.Error("Error TestArrayError2")
	}

}