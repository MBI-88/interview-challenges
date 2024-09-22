package main

import "fmt"

var (
	ecgSeq  = []int{1, 2}
	factors []int
)

func findFactors(n int) int {
	if n == 1 {
		return n
	}
	if ecgSeq[len(ecgSeq)-1]%n == 0 {
		factors = append(factors, n)
	}
	return findFactors(n - 1)
}

func checkNumber(n int) bool {
	for _, y := range ecgSeq {
		if n == y {
			return false
		}
	}
	return true
}

func generatECG(p int) {
	for i := range 3 * p {
		if i > 2 {
			findFactors(ecgSeq[len(ecgSeq)-1])
			for _, x := range factors {
				if i%x == 0 {
					if checkNumber(i) {
						ecgSeq = append(ecgSeq, i)
						factors = []int{}
						generatECG(p)
					}

				}
			}
		}
	}

}

func ecgSeqIndex(number int) int {
	var index int
	generatECG(number)
	for p, elem := range ecgSeq {
		if number == elem {
			index = p
			break
		}
	}
	return index

}

func main() {
	fmt.Printf("Input %d  index %d\n", 3, ecgSeqIndex(3))
}
