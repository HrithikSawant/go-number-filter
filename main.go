package main

import (
	"fmt"
)

func main() {
	// Input numbers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// Apply different filters using the functions defined in the same package
	evenNumbers := filterAll(numbers, isEven)
	oddNumbers := filterAll(numbers, isOdd)
	primeNumbers := filterAll(numbers, isPrime)
	oddPrimeNumbers := filterAll(numbers, isOdd, isPrime)
	evenAndMultipleOfFive := filterAll(numbers, isEven, func(n int) bool { return n%5 == 0 })
	oddAndGreaterThanTen := filterAll(numbers, isOdd, func(n int) bool { return n > 10 })

	// Output results
	fmt.Println("Even Numbers:", evenNumbers)
	fmt.Println("Odd Numbers:", oddNumbers)
	fmt.Println("Prime Numbers:", primeNumbers)
	fmt.Println("Odd Prime Numbers:", oddPrimeNumbers)
	fmt.Println("Even and Multiples of 5:", evenAndMultipleOfFive)
	fmt.Println("Odd and Greater Than 10:", oddAndGreaterThanTen)
}
