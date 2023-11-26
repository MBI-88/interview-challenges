package cmd

import (
	"log"
	"math"
	"strconv"
)

func CapturingRainwater(heights []int) int {
	var captured []int
	for index, h := range heights {
		maxLeft := 0
		maxRight := 0
		water := 0
		l := heights[:index]
		r := heights[index+1:]

		if len(l) > 0 {
			maxLeft = max(l)
		}

		if len(r) > 0 {
			maxRight = max(r)
		}

		if maxLeft != 0 && maxRight != 0 {
			water = int(math.Min(float64(maxLeft), float64(maxRight)) - float64(h))
		}

		if water > 0 {
			captured = append(captured, water)
		}

	}
	return total(captured)
}

func max(arr []int) int {
	var maxN = arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxN {
			maxN = arr[i]
		}
	}
	return maxN
}

func total(arr []int) int {
	var t int
	for _, v := range arr {
		t += v
	}
	return t
}

func StringToInt(str []string) []int {
	var result []int
	for _, s := range str {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, i)
	}
	return result

}
