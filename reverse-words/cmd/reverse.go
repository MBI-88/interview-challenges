package cmd

import (
	"fmt"
	"log"
	"strings"
)

// Invert words
func WordReverse(word string) string {
	var reversed string
	if word == "" {
		log.Fatal(fmt.Errorf("Empty string given"))
	}
	arr := strings.Fields(word)
	result, err := invertArray(arr)
	if err != nil {
		log.Fatal(err)
	}
	reversed = strings.Join(result, " ")
	return reversed
}

// Invert array order
func invertArray(arr []string) ([]string, error) {
	var newArr []string

	for i := len(arr) - 1; i >= 0; i-- {
		newArr = append(newArr, arr[i])
	}

	return newArr, nil
}

// Convert from Array to String
func ArrToString(arr []string) string {
	if len(arr) == 0 {
		log.Fatal(fmt.Errorf("Empty Array"))
	}
	return strings.Join(arr, " ")
}