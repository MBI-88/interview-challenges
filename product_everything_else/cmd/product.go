package cmd



func ProductOfOthers(arr []int) []int {
	var arrProduct []int
	
	for i := 0; i < len(arr); i++ {
		count := 1
		for x := 0; x < len(arr); x++ {
			if i != x {
				count *= arr[x]

			}
			if x == len(arr) - 1 {
				arrProduct = append(arrProduct, count)

			}
		}
	}

	return arrProduct
}