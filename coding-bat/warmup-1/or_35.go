package main

import (
	"fmt"
)

// Return true if the given non-negative number is a multiple of 3 or a multiple of 5.
// Use the % "mod" operator -- see Introduction to Mod
//
// or35(3) → true
// or35(10) → true
// or35(8) → false

func main() {
	name := "Or 35"
	fmt.Println(name, "1:", or35(3))
	fmt.Println(name, "2:", or35(10))
	fmt.Println(name, "3:", or35(8))
	fmt.Println(name, "4:", or35(-3))
}

func or35(n int) bool {

	if n >= 0 && (n % 3 == 0 || n % 5 == 0) {
		return true
	}

	return false
}

