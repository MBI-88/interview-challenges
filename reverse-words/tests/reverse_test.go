package tests

import (
	"testing"
	"words/cmd"
)




func TestReverseWords(t *testing.T) {
	sample := []struct{
		payload string
		test   string
	}{
		{
			"Codecademy rules", "rules Codecademy",
		},
		{
			"May the Fourth be with you", "you with be Fourth the May",
		},
		{
			"Reverse Words order", "order Words Reverse",
		},
	}

	for _, word := range sample {
		result := cmd.WordReverse(word.payload)
		if word.test != result {
			t.Fatalf("Error, expected %s and recived %s", word.test, result)
		}
	}
	
}