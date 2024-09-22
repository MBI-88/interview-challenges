package merge

import (
	"fmt"
)

// Merge struct Big O complexity of O(n*log(n)) worse case
type Merge struct {
	Data []int
}

// Sort method
func (m *Merge) Sort() {
	m.Data = m.handleSort(m.Data)
}

// handleSort method for  divide array
func (m Merge) handleSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	left := m.handleSort(data[:len(data)/2])
	right := m.handleSort(data[len(data)/2:])
	return m.merge(left, right)
}

// Merge method
func (m Merge) merge(left, right []int) (result []int) {
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

// Print method, print array sorted
func (m Merge) Print() {
	fmt.Printf("[")
	for p, i := range m.Data {
		if p == 0 {
			fmt.Printf("%d", i)
			continue
		}
		fmt.Printf(",%d", i)
	}
	fmt.Printf("]\n")
}
