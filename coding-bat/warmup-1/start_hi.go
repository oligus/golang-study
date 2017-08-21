package main

import (
	"fmt"
	"strings"
)

// Given a string, return true if the string starts with "hi" and false otherwise.
//
// startHi("hi there") → true
// startHi("hi") → true
// startHi("hello hi") → false

func main() {
	name := "Start hi"
	fmt.Println(name, "1:", startHi("hi there"))
	fmt.Println(name, "2:", startHi("hi"))
	fmt.Println(name, "3:", startHi("hello hi"))
}

func startHi(str string) bool {
	return strings.HasPrefix(str, "hi")
}
