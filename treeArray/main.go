package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"text/scanner"
)

func strToarry(array string) [][]int {
	var (
		s       scanner.Scanner
		dataset [][]int
	)
	s.Mode = scanner.ScanInts
	s.Init(strings.NewReader(array))
	flag := true
	for flag {
		token := s.Scan()
		switch token {
		case '[':
			continue
		case '(':
			continue
		case scanner.Int:
			child, err := strconv.Atoi(s.TokenText())
			if err != nil {
				log.Fatalln(err)
			}
			s.Scan()
			s.Scan()
			father, err := strconv.Atoi(s.TokenText())
			data := []int{child, father}
			dataset = append(dataset, data)

		case ']':
			flag = false

		case ')':
			continue

		default:
			continue

		}

	}

	return dataset
}

func findRaiz(array [][]int) int { // [[3 2] [5 2] [2 4] [1 4] [6 1] [5 1]]
	var (
		raiz  int
		found bool
	)
	for i := 0; i < len(array); i++ {
		raiz = array[i][1]
		for j := 0; j < len(array); j++ {
			if raiz == array[j][0] {
				found = false
				break
			} else if j == len(array)-1 {
				found = true
			}
		}
		if found {
			break
		}

	}

	return raiz
}

func search(father, steps int, data [][]int) int {
	for i := 0; i < len(data); i++ {
		if father == data[i][1] {
			steps++
			steps = search(data[i][0], steps, data)
		} 
	}
	return steps
}

// FindTree says if array ca be tree or not
func FindTree(array string) bool {
	var steps int 
	raiz := findRaiz(strToarry(array))
	arrayData := strToarry(array)
	steps = search(raiz, steps, arrayData)
	if steps == len(arrayData) {
		return true
	}
	return false
}

func main() {
	// (child,father)
	//arraytree1 := "[(3,2),(5,2),(2,4),(1,4),(6,1),(5,1)]"
	arraytree2 := "[(2,3),(1,3),(4,2),(5,2),(7,8)]"

	if FindTree(arraytree2) {
		fmt.Println("Tree Ok!")
	} else {
		fmt.Println("No Tree")
	}

}
