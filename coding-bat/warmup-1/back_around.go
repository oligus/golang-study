package main

import (
	"fmt"
)

// Given a string, take the last char and return a new string with the
// last char added at the front and back, so "cat" yields "tcatt".
//
// backAround("cat") → "tcatt"
// backAround("Hello") → "oHelloo"
// backAround("a") → "aaa"

func main() {
	name := "Back around"
	fmt.Println(name, "1:", backAround("cat"))
	fmt.Println(name, "2:", backAround("Hello"))
	fmt.Println(name, "3:", backAround("a"))
}

func backAround(str string) string {
	return lastChar(str) + str + lastChar(str)
}

func lastChar(str string) string {
	return string([]rune(str)[len(str) - 1])
}
