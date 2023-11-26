package tests

import (
	"testing"
	"unique/cmd"
)


func TestUniqueCharacter(t *testing.T) {
	payload := []struct{
		word string
		result bool
	}{
		{
			"apple",false,
		},
		{
			"love", true,
		},
		{
			"future", false,
		},
		{
			"freedom", false,
		},
		{
			"cutom", true,
		},

	}

	for _, item := range payload {
		if r := cmd.UniqueCharacters(item.word); r != item.result {
			t.Fatalf("Expected %t recived %t", item.result, r)
		}
	}

}