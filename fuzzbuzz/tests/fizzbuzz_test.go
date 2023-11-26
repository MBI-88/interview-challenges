package tests

import (
	"fuzz/cmd"
	"testing"
)



func TestFizzBuzz(t *testing.T) {
	payload := []struct{
		number int 
		result []interface{}
	}{
		{
			10, []interface{}{1, 2, "Fizz", 4, "Buzz", "Fizz", 7, 8, "Fizz", "Buzz"},
		},
		{
			16,[]interface{}{1, 2, "Fizz", 4, "Buzz", "Fizz", 7, 8, "Fizz", "Buzz", 11,"Fizz",13,14, "FizzBuzz", 16},
		},
		{
			20,[]interface{}{1, 2, "Fizz", 4, "Buzz", "Fizz", 7, 8, "Fizz", "Buzz", 11,"Fizz",13,14, "FizzBuzz", 16, 17,"Fizz",19,"Buzz"},
		},
	}

	for _, item := range payload {
		result := cmd.FizzBuzz(item.number)
		for i := 0; i < item.number - 1; i++ {
			if result[i] != item.result[i] {
				t.Fatalf("Expected %v recived %v number %d", result[i], item.result[i], item.number)
			}
		}
	}

}