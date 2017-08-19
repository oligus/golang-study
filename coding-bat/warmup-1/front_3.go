package main

import (
	"fmt"
)

// Given a string, we'll say that the front is the first 3 chars of the string.
// If the string length is less than 3, the front is whatever is there.
// Return a new string which is 3 copies of the front.
//
// front3("Java") → "JavJavJav"
// front3("Chocolate") → "ChoChoCho"
// front3("abc") → "abcabcabc"

func main() {
	name := "Missing char"
	fmt.Println(name, "1:", front3("Java"))
	fmt.Println(name, "2:", front3("Chocolate"))
	fmt.Println(name, "3:", front3("abc"))
	fmt.Println(name, "3:", front3("mo"))
}

func front3(str string) string {
	var front string
	var length int

	if len(str) < 3 {
		length = len(str)
	} else {
		length = 3
	}

	fmt.Println(length)

	for i := 0; i < length; i++ {
		front = front + string([]rune(str)[i])
	}
	
	return front + front + front
}


