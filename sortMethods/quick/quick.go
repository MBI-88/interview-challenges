package quick

import (
	"fmt"
)

// Quick struct Big O complexity of O(n*log(n)) worse case log(n)
type Quick struct {
	Data []int
}

// Sort method
func (q *Quick) Sort() {
	q.Data = q.handleSort(q.Data, 0, len(q.Data)-1)

}

// partitioner method divide array
func (q Quick) partitioner(data []int, low, high int) (result []int, p int) {
	pivote := data[high]
	p = low
	for j := low; j < high; j++ {
		if data[j] < pivote {
			data[p], data[j] = data[j], data[p]
			p++
		}
	}
	data[p], data[high] = data[high], data[p]
	return data, p
}

// handleSort method sort array
func (q Quick) handleSort(data []int, low, high int) []int {
	if low < high {
		arr, p := q.partitioner(data, low, high)
		data = q.handleSort(arr, low, p-1)
		data = q.handleSort(arr, p+1, high)
	}
	return data
}

// Print method, print array sorted
func (q Quick) Print() {
	fmt.Printf("[")
	for p, i := range q.Data {
		if p == 0 {
			fmt.Printf("%d", i)
			continue
		}
		fmt.Printf(",%d", i)
	}
	fmt.Printf("]\n")
}
