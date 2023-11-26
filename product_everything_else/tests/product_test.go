package tests

import (
	"product/cmd"
	"testing"
)

func TestProduct(t *testing.T) {
	data := []struct {
		source []int
		test   []int
	}{
		{
			[]int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24},
		},
		{
			[]int{5, 5, 5}, []int{25, 25, 25},
		},
		{
			[]int{6, 7, 8}, []int{56, 48, 42},
		},
	}

	for _, d := range data {
		result := cmd.ProductOfOthers(d.source)
		for i := 0; i < len(result); i++ {
			if result[i] != d.test[i] {
				t.Fatalf("Expected %#v recived %#v\n", d.test,result)
			}
		}
	}

}
