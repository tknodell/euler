package main

import "fmt"

const (
	maxNum = 100
)

func calc(n int) int {
	if n%2 == 0 { // even
		return n / 2
	} else { // odd
		return 3*n + 1
	}
}

func getChainLength(n int) int {
	var chainLength int = 1 // Include original number
	for n != 1 {
		n = calc(n)
		chainLength++
	}
	return chainLength
}

func main() {
	for i := 1; i < maxNum; i++ {
		fmt.Println(i, getChainLength(i))
	}
}
