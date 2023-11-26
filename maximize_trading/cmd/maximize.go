package cmd

import (
	"strconv"
)

// Return the best trainding profit
func MaxProfitDays(week []int) []int {
	var maxProfit []int
	if len(week) != 8 {
		return nil
	}
	tobuy := cheapest(week)
	tosell := expensive(week)
	maxProfit = append(maxProfit, tobuy,tosell)
	return maxProfit
}

// Return index of maximun
func expensive(arr []int) int {
	day := len(arr)
	max := arr[len(arr) - 1]
	for index, value := range arr {
		if value > max {
			max = value
			day = index
		}
	}
	return day
}

// Return index of minimun
func cheapest(arr []int) int {
	var day int
	min := arr[len(arr) - 1]
	for index, value := range arr {
		if value < min {
			min = value
			day = index
		}
	}
	return day
}

// Convert from string array to int array
func StringToInt( arr []string) ([]int, error) {
	var arrInt []int 
	for _, value := range arr {
		 input, err := strconv.Atoi(value)
		 if err != nil {
			continue
		 }else {
			arrInt = append(arrInt, input)
		 }
	}
	return arrInt, nil

}