package main

import (
	"fmt"
)

// Given 2 int values, return true if they are both in the range 30..40 inclusive,
// or they are both in the range 40..50 inclusive.
//
// in3050(30, 31) → true
// in3050(30, 41) → false
// in3050(40, 50) → true

func main() {
	name := "In 3050"
	fmt.Println(name, "1:", in3050(30, 31))
	fmt.Println(name, "2:", in3050(30, 41))
	fmt.Println(name, "3:", in3050(40, 50))
}

func in3050(a, b int) bool {
	if a >= 30 && b <= 40 {
		return true
	}

	if a >= 40 && b <= 50 {
		return true
	}

	return false
}

