package main

import (
	"fmt"
)

// We'll say that a number is "teen" if it is in the range 13..19 inclusive.
// Given 3 int values, return true if 1 or more of them are teen.
//
// hasTeen(13, 20, 10) → true
// hasTeen(20, 19, 10) → true
// hasTeen(20, 10, 13) → true

func main() {
	name := "Has teen"
	fmt.Println(name, "1:", hasTeen(13, 20, 10))
	fmt.Println(name, "2:", hasTeen(20, 19, 10))
	fmt.Println(name, "3:", hasTeen(20, 10, 13))
}

func hasTeen(a, b, c int) bool {
	return (a >= 13 && a <= 19) || (b >= 13 && b <= 19) || (c >= 13 && c <= 19)
}
