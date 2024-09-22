package main

import "testing"

func TestVowelSkewers(t *testing.T) {
	source := []struct {
		S string
		T bool
	}{
		{S: "B--A--N--A--N--A--S", T: true,},
		{S: "A--X--E", T: false,},
		{S: "C-L-A-P", T: false},
		{S: "M--A---T-E-S", T: false},
	}

	for _, test := range source {
		if r := isAuthenticSkewer(test.S); r != test.T {
			t.Fatalf("No match for %s result %t != %t", test.S, test.T, r)
		}
	}
}