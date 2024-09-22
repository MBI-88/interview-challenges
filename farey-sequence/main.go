package main

import (
	"fmt"
	
)

func farey(n int) []string {
	result := []string{"0/1"}

	for i := 1; i < n; i++ {
		for j := n; j > 0; j-- {
			if float32(i)/float32(j) < .50 {
				result = append(result, fmt.Sprintf("%d/%d", i,j))
			}else if (float32(i) + float32(1))/float32(n) < .50 {
				break
			}else {
				result = append(result, "1/2")
				goto step2
			}
		}
	}

	step2:
	for i := 1; i < n; i++ {
		for j := 1; j <= n; j++ {
			if float32(i)/float32(j) > .50 && float32(i)/float32(j) < 1 {
				result = append(result, fmt.Sprintf("%d/%d", i,j))
			}
		}
	}
	
	result = append(result, "1/1")

	return result
}

func main() {
	//fmt.Println(farey(5))
	//fmt.Println(farey(4))
	//fmt.Println(farey(1))
	fmt.Println(farey(10))
}
