package main

import (
	"fmt"
	"math/big"
)

func main() {
	var sum int64
	maxNum := int64(2000000)

	for i := int64(0); i < maxNum; i++ {
		if big.NewInt(i).ProbablyPrime(0) {
			sum += i
		}
	}

	fmt.Println(sum)
}
