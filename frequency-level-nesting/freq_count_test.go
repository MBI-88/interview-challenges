package main

import "testing"

func TestFreqCount(t *testing.T) {
	testcase := []struct {
		str    string
		obj    string
		output map[int]int
	}{
		{
			str: "[1,4,4,[1,1,[1,2,1,1]]]",
			obj: "1",
			output: map[int]int{
				0: 1,
				1: 2,
				2: 3,
			},
		},
		{
			str: "[1,5,5,[5,[1,2,1,1],5,5],5,[5]]",
			obj: "5",
			output: map[int]int{
				0: 3,
				1: 4,
				2: 0,
			},
		},
		{
			str: "[1,[2],1,[[2]],1,[[[2]]],1,[[[[2]]]]]",
			obj: "2",
			output: map[int]int{
				0: 0,
				1: 1,
				2: 1,
				3: 1,
				4: 1,
			},
		},
	}

	for _, p := range testcase {
		result  := freqCount(p.str, p.obj)
		for key,val := range result {
			if p.output[key] != val {
				t.Fatalf("Expected %d == %d", p.output[key], val)
			}
		}
	}

}
