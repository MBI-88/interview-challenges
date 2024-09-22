package main

import "testing"

func TestLongestAltSubstring(t *testing.T) {
	testStruct := []struct{
		Test string
		Validate string
	}{
		{
			Test: "225424272163254474441338664823",
			Validate: "272163254",
		},
		{
			Test: "594127169973391692147228678476",
			Validate: "16921472",
		},
		{
			Test: "721449827599186159274227324466",
			Validate: "7214",
		},
	}

	for _, sample := range testStruct {
		if  longestAltSubtring(sample.Test) !=  sample.Validate {
			t.Fatalf("Input %s Expected output: %s", sample.Test, sample.Validate)
		}
	}
}