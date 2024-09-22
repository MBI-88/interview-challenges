package insertion


import (
	"fmt"
)

// Insertion struct Big O complexity of O(n^2)
type Insertion struct {
	Data []int
}

// Sort method
func (in *Insertion) Sort() {
	in.Data = in.handleInsertion(in.Data)
}

// handleInsertion method to sort the array
func (in Insertion) handleInsertion(data []int) []int {
	if len(data) == 0 {return data}
	for i := 0; i < len(data); i++ {
		for j := i; j > 0 && data[j-1] > data[j]; j-- {
			data[j],data[j-1] = data[j-1],data[j]
		}
	}
	return data
}

// Print menthod print sorted array
func (in Insertion) Print(){
	fmt.Printf("[")
	for p,i := range in.Data {
		if p == 0 {
			fmt.Printf("%d",i)
			continue
		}
		fmt.Printf(",%d",i)
	}
	fmt.Printf("]\n")
}