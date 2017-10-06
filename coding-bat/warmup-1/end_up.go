package main

import (
	"fmt"
	"strings"
)

// Given a string, return a new string where the last 3 chars are now in upper case.
// If the string has less than 3 chars, uppercase whatever is there. Note that str.toUpperCase()
// returns the uppercase version of a string.
//
// endUp("Hello") → "HeLLO"
// endUp("hi there") → "hi thERE"
// endUp("hi") → "HI"

func main() {
	name := "String e"
	fmt.Println(name, "1:", endUp("Hello"))
	fmt.Println(name, "2:", endUp("hi there"))
	fmt.Println(name, "3:", endUp("hi"))
}

func endUp(str string) string {

	if len(str) < 3 {
		return strings.ToUpper(str)
	}

	last3 := strings.ToUpper(str[len(str)-3:])
	first := str[:len(str)-3]

	return first + last3
}

