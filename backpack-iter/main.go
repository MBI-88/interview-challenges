package main

import (
	"fmt"
)


type Item struct {
	Value  int
	Weight int
}

func knapSack(W int, items []Item) int {
	n := len(items)
	
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, W+1)
	}

	for i := 0; i <= n; i++ {
		for w := 0; w <= W; w++ {
			if i == 0 || w == 0 {
				dp[i][w] = 0
			} else if items[i-1].Weight <= w {
				valorConItem := items[i-1].Value + dp[i-1][w-items[i-1].Weight]
				if valorConItem > dp[i-1][w] {
					dp[i][w] = valorConItem
				} else {
					dp[i][w] = dp[i-1][w]
				}
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
		fmt.Printf("%v\n", dp)
	}
	return dp[n][W]
}

func main() {
	items := []Item{
		{Value: 300, Weight: 10},
		{Value: 100, Weight: 5},
		{Value: 200, Weight: 30},
		{Value: 120, Weight: 3},
		{Value: 521, Weight: 30},
	}

	W := 30

	resultado := knapSack(W, items)
	fmt.Printf("Max value gotton: %d\n", resultado)
}