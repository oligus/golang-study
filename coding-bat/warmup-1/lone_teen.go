package main

import (
	"fmt"
)

// We'll say that a number is "teen" if it is in the range 13..19 inclusive.
// Given 2 int values, return true if one or the other is teen, but not both.
//
// loneTeen(13, 99) → true
// loneTeen(21, 19) → true
// loneTeen(13, 13) → false
// loneTeen(99, 99) → false

func main() {
	name := "Has teen"
	fmt.Println(name, "1:", loneTeen(13, 99))
	fmt.Println(name, "2:", loneTeen(21, 19))
	fmt.Println(name, "3:", loneTeen(13, 13))
	fmt.Println(name, "4:", loneTeen(99, 99))
}

func loneTeen(a, b int) bool {
	return (isTeen(a) || isTeen(b)) && !(isTeen(a) && isTeen(b))
}

func isTeen(age int) bool {
	return age >= 13 && age <= 19
}
