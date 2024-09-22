package main

import "testing"

func TestFarey(t *testing.T) {
	source := []struct{
		Input int
		Ouput []string
	}{
		{
			5,[]string{"0/1", "1/5", "1/4", "1/3", "2/5", "1/2", "2/3", "3/4", "3/5", "4/5", "1/1"},
		},
		{
			1, []string{"0/1","1/1"},
		},
		{
			4, []string{"0/1","1/4","1/3","1/2","2/3","3/4","1/1"},
		},
	}

	for _, s := range source {
		result := farey(s.Input)
		for i := 0; i < len(result); i++ {
			if result[i] != s.Ouput[i] {
				t.Fatalf("Exection %s != %s", result[i],s.Ouput[i])
			}
		}
	}
}