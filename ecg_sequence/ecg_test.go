package main

import "testing"

func TestEcgSeq(t *testing.T) {
	sample := []struct{
		Test int
		Val int
	}{
		{Test: 3, Val: 4},
		{Test: 5, Val: 9},
		{Test: 7, Val: 13},
		{Test: 8, Val: 7},
	}

	for _, smp := range sample {
		if ecgSeqIndex(smp.Test) != smp.Val {
			t.Fatalf("Input %d expected %d", smp.Test, smp.Val)
		}
	}
}