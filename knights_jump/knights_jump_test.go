package main

import "testing"

func TestKnightJump(t *testing.T) {
	testSample := []struct{
		Test string
		Validate string
	}{
		{
			Test: "F4",
			Validate: "D3,D5,E2,E6,G2,G6,H3,H5",
		},
		{
			Test: "A1",
			Validate: "B3,C2",

		},
		{
			Test: "E2",
			Validate: "C1,C3,D4,F4,G1,G3",
		},
	}

	for _, ts := range testSample {
		if knightsJump(ts.Test) != ts.Validate {
			t.Fatalf("Input %s Expected output: %s", ts.Test, ts.Validate)
		}
	}

}