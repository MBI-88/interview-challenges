package cmd

import (
	"strconv"
	"strings"
)


// Solve knapsack problem
func Knapsack(weights, values []int, unit int) int {
	var (
		backpack  []int
		total int
	)

	for i := 0; i < len(weights); i++ {
		if r := isMaximun(values[i],unit, values, weights); r {
			if sum(&total, weights[i], unit) {
				backpack = append(backpack, values[i])
			}
		}
		for x := i + 1; x < len(weights); x++ {
			if r := isMaximun(values[i]+values[x],unit, values, weights); r {
				if len(backpack) == 0 {
					if sum(&total, weights[x]+weights[i], unit) {
						backpack = append(backpack, values[x]+values[i])
					}
				}else {
					if sum(&total, weights[x], unit) {
						backpack = append(backpack, values[x])
					}
				}

			}
		}
	}

	return getTotal(backpack)
}

// Is maximun return boolean
func isMaximun(total,u int, v, w []int) bool {
	for i := 0; i < len(v); i++ {
		if total < v[i] {
			return false
		}
		for x := i + 1; x < len(v); x++ {
			if total < v[i] + v[x] && w[i] + w[x] <= u {
				return false
			}
		}

	}
	return true
}

// Save total added
func sum(t *int, w, u int) bool {
	if *t+w > u {
		return false
	}
	*t += w
	return true
}

// Return the maximun recovered value
func getTotal(arr []int) int {
	var temp int
	for _, v := range arr {
		temp += v
	}
	return temp
}

// Convert string to int
func StrToInt(arr string) []int {
	var val []int
	strArr := strings.Split(arr, ",")
	for _, str := range strArr {
		v, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		val = append(val, v)
	}

	return val
}
