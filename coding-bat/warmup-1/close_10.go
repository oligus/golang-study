package main

import (
	"fmt"
	"math"
)

// Given 2 int values, return whichever value is nearest to the value 10, or return 0 in the event of a tie.
// Note that Math.abs(n) returns the absolute value of a number.
//
// close10(8, 13) → 8
// close10(13, 8) → 8
// close10(13, 7) → 0

func main() {
	name := "Close 10"
	fmt.Println(name, "1:", close10(8, 13))
	fmt.Println(name, "2:", close10(13, 8))
	fmt.Println(name, "3:", close10(13, 7))
}

func close10(a, b int) int {
	resultA := math.Abs(10 - float64(a))
	resultB := math.Abs(10 - float64(b))

	if resultA == resultB {
		return 0
	}

	if resultA > resultB {
		return b
	}


	return a
}

