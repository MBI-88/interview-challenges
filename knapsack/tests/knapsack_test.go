package tests

import (
	"knap/cmd"
	"testing"
)

func TestKnapsack(t *testing.T) {
	source := []struct{
		weights []int
		values []int 
		limit int 
		result int
	}{
		{
			[]int{3,5,8},[]int{50,100,100},10,150,
		},
		{
			[]int{2,5,8},[]int{50,100,100},10,150,
		},
		{
			[]int{6,5,8},[]int{50,100,100},10,100,
		},
		{
			[]int{3,5,8,9},[]int{50,100,100,500},10,500,
		},
		{
			[]int{3,5,8,4,10},[]int{50,100,100,200,40},10,300,
		},
		{
			[]int{1,1,5,8}, []int{50,100,60,100},10,200,
		},

	}

	for _, ts := range source {
		rs := cmd.Knapsack(ts.weights, ts.values, ts.limit)
		if rs != ts.result {
			t.Fatalf("Recived %d but expected %d ", rs, ts.result)
		}
	}
	
}