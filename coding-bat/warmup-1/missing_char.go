package main

import (
	"fmt"
)

// Given a non-empty string and an int n, return a new string where the char at index n has been removed.
// The value of n will be a valid index of a char in the original string
// (i.e. n will be in the range 0..str.length()-1 inclusive).
//
// missingChar("kitten", 1) → "ktten"
// missingChar("kitten", 0) → "itten"
// missingChar("kitten", 4) → "kittn"

func main() {
	name := "Missing char"
	fmt.Println(name, "1:", missingChar("kitten", 1))
	fmt.Println(name, "2:", missingChar("kitten", 0))
	fmt.Println(name, "3:", missingChar("kitten", 4))
}

func missingChar(str string, n int) string {
	var result string

	for i := 0; i < len(str); i++ {
		if n != i {
			result = result + string([]rune(str)[i])
		}
	}

	return result
}


