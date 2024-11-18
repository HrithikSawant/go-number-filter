package main

type Condition func(n int) bool

func isEven(number int) bool {
	return number%2 == 0
}

func isOdd(number int) bool {
	return number%2 != 0
}

func isPrime(number int) bool {
	if number <= 1 {
		return false // 0 and 1 are not prime
	}

	for divisor := 2; divisor*divisor <= number; divisor++ { // Check up to √number
		if number%divisor == 0 {
			return false // Found a factor, not prime
		}
	}
	return true
}

func evenNumbers(numbers []int) (evenNums []int) {
	for _, number := range numbers {
		if isEven(number) {
			evenNums = append(evenNums, number)
		}
	}
	return evenNums
}

func oddNumbers(numbers []int) (oddNums []int) {
	for _, number := range numbers {
		if isOdd(number) {
			oddNums = append(oddNums, number)
		}
	}
	return oddNums
}

func primeNumbers(numbers []int) (primeNums []int) {
	for _, number := range numbers {
		if isPrime(number) {
			primeNums = append(primeNums, number)
		}
	}
	return primeNums
}

func oddPrimeNumbers(numbers []int) (oddPrimeNums []int) {
	for _, number := range numbers {
		if isPrime(number) && isOdd(number) {
			oddPrimeNums = append(oddPrimeNums, number)
		}
	}
	return oddPrimeNums
}

func filterAll(numbers []int, conditions ...Condition) (filteredNumbers []int) {
	for _, number := range numbers {
		matchesAllConditions := true
		for _, condition := range conditions {
			if !condition(number) {
				matchesAllConditions = false
			}
		}

		if matchesAllConditions {
			filteredNumbers = append(filteredNumbers, number)
		}
	}
	return filteredNumbers
}

func filterAny(numbers []int, conditions ...Condition) (filteredNumbers []int) {
	for _, number := range numbers {
		matchesAnyConditions := false
		for _, condition := range conditions {
			if condition(number) {
				matchesAnyConditions = true
				break
			}
		}

		if matchesAnyConditions {
			filteredNumbers = append(filteredNumbers, number)
		}
	}
	return filteredNumbers
}
