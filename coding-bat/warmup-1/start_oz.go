package main

import (
	"fmt"
)

// Given a string, return a string made of the first 2 chars (if present),
// however include first char only if it is 'o' and include the second only if it is 'z',
// so "ozymandias" yields "oz".
//
// startOz("ozymandias") → "oz"
// startOz("bzoo") → "z"
// startOz("oxx") → "o"

func main() {
	name := "Start oz"
	fmt.Println(name, "1:", startOz("ozymandias"))
	fmt.Println(name, "2:", startOz("bzoo"))
	fmt.Println(name, "3:", startOz("oxx"))
}

func startOz(str string) string {
	firstChar := string([]rune(str)[0])
	secondChar := string([]rune(str)[1])
	result := ""

	if firstChar == "o" {
		result = result + firstChar
	}

	if secondChar == "z" {
		result = result + secondChar
	}

	return result
}
