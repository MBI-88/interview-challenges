package cmd




func UniqueCharacters(word string) bool{
	result := mapUnique(word)
	
	for _, i := range word {
		if result[i] > 1 {
			return false
		}
	}

	return true
}

// Return a map of runes counted
func mapUnique(w string) map[rune]int {
	m := make(map[rune]int)

	for _, i := range w {
		m[i]++
	}
	return m
}