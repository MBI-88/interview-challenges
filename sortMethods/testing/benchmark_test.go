package testing

import (
	"math/rand"
	"methods/bubble"
	"methods/merge"
	"methods/quick"
	"methods/insertion"
	"testing"
)

func generateArray() []int {
	array := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		array = append(array, rand.Int())
	}
	return array
}

func BenchmarkBubble(b *testing.B) {
	array := generateArray()
	bub := bubble.Bubble{Data: array}
	for i := 0; i < b.N; i++ {
		bub.Sort()
	}

}

func BenchmarkMerge(b *testing.B) {
	array := generateArray()
	mer := merge.Merge{Data: array}
	for i := 0; i < b.N; i++ {
		mer.Sort()
	}

}

func BenchmarkQuick(b *testing.B) {
	array := generateArray()
	qu := quick.Quick{Data: array}
	for i := 0; i < b.N; i++ {
		qu.Sort()
	}
}

func BenchmarkInsertion(b *testing.B) {
	array := generateArray()
	in := insertion.Insertion{Data:array}
	for i := 0; i < b.N; i++ {
		in.Sort()
	}
}