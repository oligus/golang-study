package main

import (
	"fmt"
)

// Given a string, return a new string where the first and last chars have been exchanged.
//
// frontBack("code") → "eodc"
// frontBack("a") → "a"
// frontBack("ab") → "ba"

func main() {
	name := "Front back"
	fmt.Println(name, "1:", frontBack("code"))
	fmt.Println(name, "2:", frontBack("a"))
	fmt.Println(name, "3:", frontBack("ab"))
}

func frontBack(str string) string {
	var result string

	for i := len(str) - 1; i >= 0; i-- {
		result = result + string([]rune(str)[i])
	}
	
	return result
}


