package tests

import (
	"fmt"
	"rainwater/cmd"
	"testing"
)

func TestRainWater(t *testing.T) {
	sample := []struct {
		source []int
		test   int
	}{
		{
			[]int{4, 2, 1, 3, 0, 1, 2}, 6,
		},
		{
			[]int{1,2,1}, 0,
		},
		{
			[]int{1,3,0,5}, 3,
		},
	}

	for _, v := range sample {
		resutl := cmd.CapturingRainwater(v.source)
		if resutl != v.test {
			t.Fatal(fmt.Errorf("Error, expected %d recived %d", v.test, resutl))
		}
	}

}
