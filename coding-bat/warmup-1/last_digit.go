package main

import (
	"fmt"
)

// Given two non-negative int values, return true if they have the same last digit, such as with 27 and 57.
// Note that the % "mod" operator computes remainders, so 17 % 10 is 7.
//
// lastDigit(7, 17) → true
// lastDigit(6, 17) → false
// lastDigit(3, 113) → true

func main() {
	name := "String e"
	fmt.Println(name, "1:", lastDigit(7, 17))
	fmt.Println(name, "2:", lastDigit(6, 17))
	fmt.Println(name, "3:", lastDigit(3, 113))
}

func lastDigit(a, b int) bool {
	remain := b % 10

	if remain == a {
		return true
	}

	return false
}

