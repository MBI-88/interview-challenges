package main

import (
	"testing"
)

func TestReversedIndex(t *testing.T) {

	source := []struct{
		input []any 
		status string
		output []any
	}{
		{
			input: []any{2, 4, 6, 8, 10},status:"odd", output: []any{2, 6, 10},
		},
		{input: []any{"E", "D", "A", "B", "I", "T"}, status: "even", output: []any{"E", "A", "I"} },
		{input: []any{")", "(", "*", "&", "^", "%", "$", "#", "@", "!"},status: "even", output: []any{")", "*", "^", "$", "@"} },
		{input: []any{"A", "R", "B", "I", "T", "R", "A", "R", "I", "L", "Y"}, status: "even", output: []any{"R", "I", "R", "R", "L"}},
	}

	for _, obj := range source {
		temp := getItemsAt(obj.input, obj.status)
		for i := 0; i < len(obj.output); i++ {
			if temp[i] != obj.output[i] {
				t.Fatal("Input and output are different")
			}
		}
	}
}