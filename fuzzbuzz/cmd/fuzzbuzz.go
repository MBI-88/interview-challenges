package cmd 

// FizzBuzz function 
func FizzBuzz(n int) []interface{} {
	if n <= 0 {
		return nil
	}
	var array []interface{} 

	arrInt := makeRange(n)

	for _, item := range arrInt {
		isFizz := fizz(item)
		isBuzz := bizz(item)
		if isFizz == "Fizz" && isBuzz == "Buzz"{
			array = append(array, isFizz+isBuzz )

		}else if isFizz == "Fizz" && isBuzz != "Buzz" {
			array = append(array, isFizz)
			
		}else if isFizz != "Fizz" && isBuzz == "Buzz" {
			array = append(array, isBuzz)

		}else	{
			array = append(array, item)
		
		}
	}

	return array
}

// If n is multiple of three return Fizz
func fizz(n int) string {
	if n % 3 != 0 {
		return "no"
	} 
	return "Fizz"
}

// If n is multiple of five return Buzz
func bizz(n int) string {
	if n % 5 != 0 {
		return "no"
	}
	return "Buzz"
}

// Make a range given a number
func makeRange(n int) []int {
	var arr []int
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}
	return arr
}