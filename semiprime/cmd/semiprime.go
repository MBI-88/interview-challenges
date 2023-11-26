package cmd


func SemiprimeCount(n int) int {
	var count int
	var arr []int

	for x := 1; x < n; x++ {
		arr = append(arr, x)
	}
	
	for i := 0; i < len(arr); i++ {
		for s := 0; s < len(arr[:i]); s++ {
			if  arr[s] != 1  && arr[i] % arr[s] == 0 {
				con1 := checkprimme(arr[s])
				r := arr[i] / arr[s] 
				cons2 := checkprimme(r)

				if con1 && cons2 {
					count++
					break
				}
			}
			
		}
	}


	return count 
}

func checkprimme(n int) bool {
	for i := 2; i < n; i++ {
		if  n % i == 0 {
			return false
		}
	}
	return true
}