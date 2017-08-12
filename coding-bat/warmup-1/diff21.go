package main

import "fmt"

// Given an int n, return the absolute difference between n and 21,
// except return double the absolute difference if n is over 21.
//
// diff21(19) → 2
// diff21(10) → 11
// diff21(21) → 0

func main() {
	name := "Diff 21"
	fmt.Println(name, "1:", diff21(19))
	fmt.Println(name, "2:", diff21(10))
	fmt.Println(name, "3:", diff21(21))
}

func diff21(n int) int {
	if n > 21 {
		return (n - 21) * 2
	}

	return 21 - n
}


