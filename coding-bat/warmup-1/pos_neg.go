package main

import "fmt"

// Given 2 int values, return true if one is negative and one is positive.
// Except if the parameter "negative" is true, then return true only if both are negative.
//
// posNeg(1, -1, false) → true
// posNeg(-1, 1, false) → true
// posNeg(-4, -5, true) → true

func main() {
	name := "Pos neg"
	fmt.Println(name, "1:", posNeg(1, -1, false))
	fmt.Println(name, "2:", posNeg(-1, 1, false))
	fmt.Println(name, "3:", posNeg(-4, -5, true))
}

func posNeg(a, b int, negative bool) bool {

	if negative {
		return a < 0 && b < 0
	}

	return (a < 0 && b >= 0) || (b < 0 && a >= 0)
}


