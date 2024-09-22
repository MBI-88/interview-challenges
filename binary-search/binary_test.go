package main

import (
	"io"
	"os"
	"testing"
)

func TestBinarySearchFound(t *testing.T) {
	var arr []int 
	for item := range 100 {
		arr = append(arr, item)
	}

	output := os.Stdout
	r,w, _ := os.Pipe()
	os.Stdout = w

	binarySearch(arr, 20)
	w.Close() 
	os.Stdout = output
	temp, _ := io.ReadAll(r)

	srtOutput := string(temp)
	if srtOutput == "Not found\n" {
		t.Fatalf("Expected match")
	}
}

func TestBinarySearchNotfound(t *testing.T) {
	var arr []int 
	for item := range 50 {
		arr = append(arr, item)
	}

	output := os.Stdout
	r,w, _ := os.Pipe()
	os.Stdout = w

	binarySearch(arr, 60)
	w.Close() 
	os.Stdout = output
	temp, _ := io.ReadAll(r)
	srtOutput := string(temp)
	if srtOutput != "Not found\n" {
		t.Fatalf("Expected Not found")
	}
}