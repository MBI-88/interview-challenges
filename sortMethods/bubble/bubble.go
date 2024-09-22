package bubble

import (
	"fmt"
)

// Bubble struct Big O complexity of  O(n^2)
type Bubble struct {
	Data []int
}

// Sort method, using Bubble sort
func (b *Bubble) Sort() {
	datalen := b.lenbubble()
	for i := 0; i < datalen; i++ {
		for j := 0; j < datalen-i; j++ {
			if b.Data[j] > b.Data[j+1] {
				temp := b.Data[j+1]
				b.Data[j+1] = b.Data[j]
				b.Data[j] = temp

			}
		}
	}

}

func (b *Bubble) lenbubble() int {
	return len(b.Data) - 1
}

// Print method print bubble data
func (b *Bubble) Print() {
	fmt.Printf("[")
	for p, i := range b.Data {
		if p == 0 {
			fmt.Printf("%d", i)
			continue
		}
		fmt.Printf(",%d", i)
	}
	fmt.Printf("]\n")
}
