package main

import (
	"fmt"
)

// Given 2 positive int values, return the larger value that is in the range 10..20 inclusive,
// or return 0 if neither is in that range.
//
// max1020(11, 19) → 19
// max1020(19, 11) → 19
// max1020(11, 9) → 11

func main() {
	name := "Max 1020"
	fmt.Println(name, "1:", max1020(11, 19))
	fmt.Println(name, "2:", max1020(19, 11))
	fmt.Println(name, "3:", max1020(11, 9))
}

func max1020(a, b int) int {
	if a >= 10 && b <= 20 {
		if a > b {
			return a
		}

		return b
	}

	return 0
}

