package tests

import (
	"prime-factors/cmd"
	"testing"
)

func TestSumPrimeFactors(t *testing.T) {
	payload := []struct {
		number int
		result int
	}{
		{
			91, 20,
		},
		{
			100, 7,
		},
		{
			30, 10,
		},
		{
			21, 10,
		},
	}

	for _, item := range payload {
		result := cmd.SumePrimeFactors(item.number)
		if result != item.result {
			t.Fatalf("Expected %d recived %d", item.result, result)
		}
	}
}
