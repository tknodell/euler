package main

import "fmt"

const (
	startNum = 13
	maxNum   = 10
)

func calc(n int) int {
	if n%2 == 0 { // even
		return n / 2
	} else { // odd
		return 3*n + 1
	}
}

func main() {
	var res int = startNum

	for i := startNum; i < startNum+maxNum; i++ {
		fmt.Println(res)
		res = calc(res)
	}
}
