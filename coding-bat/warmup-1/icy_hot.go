package main

import (
	"fmt"
)

// Given two temperatures, return true if one is less than 0 and the other is greater than 100.
//
// icyHot(120, -1) → true
// icyHot(-1, 120) → true
// icyHot(2, 120) → false

func main() {
	name := "Icy hot"
	fmt.Println(name, "1:", icyHot(120, -1))
	fmt.Println(name, "2:", icyHot(-1, 120))
	fmt.Println(name, "3:", icyHot(2, 120))
}

func icyHot(temp1, temp2 int) bool {
	return (temp1 < 0 && temp2 > 100) || (temp2 < 0 && temp1 > 100)
}
