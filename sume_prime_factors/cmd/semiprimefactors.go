package cmd

func SumePrimeFactors(number int) int {
	var (
		primefactors []int
		sume         int
	)
	primefactors = checkFactors(number)

	for _, n := range primefactors {
		sume += n
	}

	return sume
}

func checkFactors(number int) []int {
	var arrN []int
	for i := 2; i < number; i++ {
		if number % i == 0 {
			if isPrime(i) {
				arrN = append(arrN, i)
			}
		}

	}
	return arrN
}

func isPrime(number int) bool {
	for i := 2; i < number; i++ {
		if number % i == 0 {
			return false
		}

	}
	return true
}
