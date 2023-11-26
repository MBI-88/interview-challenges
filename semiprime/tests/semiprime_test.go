package tests

import (
	"semiprime/cmd"
	"testing"
)

func TestSemiprine(t *testing.T) {
	var test = []struct {
		input int
		ouput int
	}{
		{
			10, 3,
		},
		{
			15, 5,
		},
		{
			20, 6,
		},
		{
			25, 8,
		},
	}

	for _, st := range test {

		if cmd.SemiprimeCount(st.input) != st.ouput {
			t.Fatalf("Error in the function recived %d expected %d", cmd.SemiprimeCount(st.input), st.ouput)
		}

	}

}
