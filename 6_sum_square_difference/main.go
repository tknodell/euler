package main

import (
	"fmt"
	"math"
)

// The sum of the squares of the first ten natural numbers is,

// 12+22+...+102=385
// The square of the sum of the first ten natural numbers is,

// (1+2+...+10)2=552=3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025−385=2640.

// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

func main() {
	fmt.Printf("%f", difference(100))
}

func difference(n float64) float64 {
	return sumOfTheSquare(n) - sumOfSquares(n)
}

func sumOfSquares(r float64) float64 {
	var sum float64
	for i := float64(1); i <= r; i++ {
		sum = sum + math.Pow(i, 2)
	}
	return sum
}

func sumOfTheSquare(r float64) float64 {
	var sum float64
	for i := float64(1); i <= r; i++ {
		sum = sum + i
	}
	return math.Pow(sum, 2)
}
