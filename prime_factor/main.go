// 1) While n is divisible by 2, print 2 and divide n by 2.
// 2) After step 1, n must be odd. Now start a loop from i = 3 to square root of n. While i divides n, print i and divide n by i. After i fails to divide n, increment i by 2 and continue.
// 3) If n is a prime number and is greater than 2, then n will not become 1 by above two steps. So print n if it is greater than 2.

package main

import (
	"fmt"
	"math"
)

func main() {
	// findFactor(13195)
	findFactor(600851475143)
}

func findFactor(n int) {
	// Step 1
	for n%2 == 0 {
		fmt.Println(n)
		n = n / 2
	}

	// Step 2
	sqrt := int(math.Sqrt(float64(n)))

	for i := 3; i <= sqrt; {
		if n%i == 0 {
			fmt.Println(i)
			n = n / i
		} else {
			i = i + 2
		}
	}

	// Step 3
	if n > 2 {
		fmt.Println(n)
	}

}
