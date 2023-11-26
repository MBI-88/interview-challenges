package test

import (
	"testing"
	"xth-number/cmd"
)


func TestFindXthNumber(t *testing.T) {
	payload := []struct {
		number int
		array []int 
		result int
	}{
		{
			2, []int{5,10,-3,-3,7,9}, -3,
		},
		{
			4, []int{5,10,-3,-3,7,9}, 7,
		},
		{
			5, []int{2,5,6,-8, 6,7,-10}, 6,
		},
	}

	for _, item := range payload {
		result := cmd.GetX(item.number, item.array)
		if result != item.result {
			t.Fatalf("Expected %d recived %d", item.result, result)
		}
	}
}