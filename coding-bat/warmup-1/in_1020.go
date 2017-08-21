package main

import (
	"fmt"
)

// Given 2 int values, return true if either of them is in the range 10..20 inclusive.
//
// in1020(12, 99) → true
// in1020(21, 12) → true
// in1020(8, 99) → false

func main() {
	name := "In 1020"
	fmt.Println(name, "1:", in1020(12, 99))
	fmt.Println(name, "2:", in1020(21, 12))
	fmt.Println(name, "3:", in1020(8, 99))
}

func in1020(a, b int) bool {
	return (a >= 10 && a <= 20) || (b >= 10 && b <= 20)
}
