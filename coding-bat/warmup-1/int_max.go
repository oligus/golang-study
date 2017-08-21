package main

import (
	"fmt"
)

// Given three int values, a b c, return the largest.
//
// intMax(1, 2, 3) → 3
// intMax(1, 2, 3) → 3
// intMax(3, 2, 1) → 3

func main() {
	name := "Int max"
	fmt.Println(name, "1:", intMax(1, 2, 3))
	fmt.Println(name, "2:", intMax(1, 2, 3))
	fmt.Println(name, "3:", intMax(3, 2, 1))
}

func intMax(a, b, c int) int {
	max := a

	if b > max {
		max = b
	}

	if c > max {
		max = c
	}

	return max
}

