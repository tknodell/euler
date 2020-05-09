package main

import "fmt"

// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

// What is the 10 001st prime number?

func main() {
	fmt.Println(countPrime(10001))
}

func countPrime(n int) int {
	var counter, primeCount int
	for primeCount <= n {
		counter++
		if isPrime(counter) {
			primeCount++
		}
		if primeCount == n {
			return counter
		}
	}
	return -1
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n-1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
