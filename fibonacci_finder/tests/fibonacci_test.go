package tests

import (
	"fibonacci/cmd"
	"testing"
)

func TestFibonacciFinder(t *testing.T) {

	data := []struct{
		k int 
		n int
	}{
		{
			3, 1,
		},
		{
			6, 5,
		},
		{
			2,1,
		},
		{
			1, 0,
		},
		{
			5, 3,
		},
	}

	for _, item := range data {
		result := cmd.FiboFinder(item.k)
		if result != item.n {
			t.Fatalf("Expetect result == item.n but recived %d != %d", result, item.n)
		}
	}
}