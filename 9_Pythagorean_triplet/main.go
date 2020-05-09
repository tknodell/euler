package main

import "fmt"

type Triplet [3]int

func main() {
	max := 500
	target := 1000

	triplets := Range(0, max)
	for t := 0; t < len(triplets); t++ {
		triplet := triplets[t]
		sum := triplet.sumTriplet()
		if sum == target {
			fmt.Println("FOUND")
			fmt.Println(triplet)
			fmt.Println(triplet.productTriplets())
			break
		}
	}
}

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	res := []Triplet{}
	for a := 0; a < max; a++ {
		for b := 0; b < max; b++ {
			for c := 0; c < max; c++ {
				switch {
				case (a*a)+(b*b) == c*c:
					res = append(res, Triplet{a, b, c})
					break
				case a*a+b*b < c*c:
					break
				}
			}
		}
	}
	return res
}

func (t Triplet) sumTriplet() int {
	return t[0] + t[1] + t[2]
}

func (t Triplet) productTriplets() int {
	return t[0] * t[1] * t[2]
}
