package main

import (
	"fmt"
	"strconv"
)

// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

// Find the largest palindrome made from the product of two 3-digit numbers.

func main() {
	result := 0
	for i := 0; i <= 999; i++ {
		for j := 0; j <= 999; j++ {
			t := i * j
			if isPalindrome(t) && t > result {
				// fmt.Println(i, j)
				result = t
			}
		}
	}
	fmt.Println(result)
}

func isPalindrome(n int) bool {
	return strconv.Itoa(n) == revString(strconv.Itoa(n))
}

func revString(n string) string {
	var result string
	for i := 0; i < len(n); i++ {
		result += (string(n[len(n)-i-1]))
	}

	return result
}
