package main

import (
	"fmt"
	"time"
)

func fillMap(word string) map[rune]int {
	mp := make(map[rune]int)
	for _, letter := range word {
		mp[letter] += 1
	}
	return mp
}

func checkAnagrams(m map[rune]int, word string) bool {
	for _, letter := range word {
		if _, exist := m[letter]; exist {
			m[letter] -= 1

		} else {
			return false
		}
	}
	for _, val := range m {
		if val != 0 {
			return false
		}
	}
	return true
}

func Anagrams(source []string) [][]string {
	var anagrams [][]string
	for i := 0; i < len(source); i++ {
		for j := i + 1; j < len(source); j++ {
			if len(source[i]) == len(source[j]) && checkAnagrams(fillMap(source[i]), source[j]) {
				anagrams = append(anagrams, []string{source[i], source[j]})
			}

		}
	}
	return anagrams
}

func main() {
	source1 := []string{"amor", "roma", "calor", "frio", "comer", "mora"}
	socurce2 := []string{"riesgo", "peligro", "amar", "solucion", "sergio", "rama"}
	source3 := []string{"alma", "demonio", "lama", "nabo", "bona", "mar", "ram"}

	for _, sr := range [][]string{source1, socurce2, source3} {
		startTime := time.Now()
		fmt.Printf("Anagrams found %s\n", Anagrams(sr))
		fmt.Printf("Time taken %d ms\n", time.Now().Sub(startTime).Microseconds())
	}

}
