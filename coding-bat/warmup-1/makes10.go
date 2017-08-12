package main

import "fmt"

// Given 2 ints, a and b, return true if one if them is 10 or if their sum is 10.
//
// makes10(9, 10) → true
// makes10(9, 9) → false
// makes10(1, 9) → true

func main() {
	name := "Makes 10"
	fmt.Println(name, "1:", makes10(9, 10))
	fmt.Println(name, "2:", makes10(9, 9))
	fmt.Println(name, "3:", makes10(1, 9))
}

func makes10(a, b int) bool {
	if a == 10 || b == 10 || (a + b) == 10 {
		return true
	}

	return false
}




