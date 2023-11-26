package tests

import (
	"maximize/cmd"
	"testing"
)




func TestMaximize(t *testing.T) {

	data := []struct{
		arr []int
		result []int
	}{
		{
			[]int{17,11,60,25,150,31,120,30},[]int{1,4},
		},
		{
			[]int{5,20,30,100,45,10,90,50},[]int{0,3},
		},
		{
			[]int{150,50,100,30,80,45,10,20},[]int{6,0},
		},
	}

	for _, item := range data {
		result := cmd.MaxProfitDays(item.arr)
		if result[0] != item.result[0] || result[1] != item.result[1] {
			t.Fatalf("Error expected result == item.result")
		}
	}
}