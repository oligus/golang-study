package main

import (
	"fmt"
)

// Given a non-empty string and an int N, return the string made starting with char 0,
// and then every Nth char of the string. So if N is 3, use char 0, 3, 6, ... and so on. N is 1 or more.
//
// everyNth("Miracle", 2) → "Mrce"
// everyNth("abcdefg", 2) → "aceg"
// everyNth("abcdefg", 3) → "adg"


func main() {
	name := "String e"
	fmt.Println(name, "1:", everyNth("Miracle", 2))
	fmt.Println(name, "2:", everyNth("abcdefg", 2))
	fmt.Println(name, "3:", everyNth("abcdefg", 3))
}

func everyNth(str string, nth int) string {
	var result string

	for i := 0; i < len(str); i++ {
		if i % nth == 0 {
			result = result + str[i:i+1]
		}
	}

	return result
}

