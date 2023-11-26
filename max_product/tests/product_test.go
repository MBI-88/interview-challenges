package tests

import (
	"maxproduct/cmd"
	"testing"
)

func TestMaximunProductFinder(t *testing.T) {

	data := []struct {
		input  []int
		k      int
		output int
	}{
		{
			[]int{-8, 6, -7, 3, 2, 1, -9}, 2, 72,
		},
		{
			[]int{9, 5, -2, -1, -10, 6}, 3, 270,
		},
		{
			[]int{5,-5,-6,3,2},2,30,
		},
		{
			[]int{-9,8,6,2,4,-9},3, 648,
		},
	}

	for _, item := range data {
		result := cmd.MaxProductFinder(item.input, item.k)
		if result != item.output {
			t.Fatalf("Expected result == item.ouput but recived %d != %d",result, item.output)
		}
	}

}
