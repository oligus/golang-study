package main

import "fmt"

// Given two int values, return their sum. Unless the two values are the same, then return double their sum.
//
// sumDouble(1, 2) → 3
// sumDouble(3, 2) → 5
// sumDouble(2, 2) → 8

func main() {
	fmt.Println("Sum Double 1:", sumDouble(1, 2))
	fmt.Println("Sum Double 2:", sumDouble(3, 2))
	fmt.Println("Sum Double 3:", sumDouble(2, 2))
}

func sumDouble(a, b int) int {
	if a == b  {
		return (a + b) * 2
	} else {
		return a + b
	}
}
