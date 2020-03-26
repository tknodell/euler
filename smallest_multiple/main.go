package main

import "fmt"

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

func main() {
	res := 1
	for isDivisibleBy(res, 20) == false {
		res++
	}
	fmt.Println(res)
}

func isDivisibleBy(n, count int) bool {
	var isDiv bool
	for i := 1; i <= count; i++ {
		if n%i == 0 {
			isDiv = true
		} else {
			return false
		}
	}
	return isDiv
}
