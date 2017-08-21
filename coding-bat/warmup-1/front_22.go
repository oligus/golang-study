package main

import (
	"fmt"
)

// Given a string, take the first 2 chars and return the string with the 2 chars added at both the front and back,
// so "kitten" yields "kikittenki". If the string length is less than 2, use whatever chars are there.
//
// front22("kitten") → "kikittenki"
// front22("Ha") → "HaHaHa"
// front22("abc") → "ababcab"

func main() {
	name := "Front 22"
	fmt.Println(name, "1:", front22("kitten"))
	fmt.Println(name, "2:", front22("Ha"))
	fmt.Println(name, "3:", front22("abc"))
}

func front22(str string)string {
	length := 2

	if len(str) < 2 {
		length = len(str)
	}

	var front string

	for i := 0; i < length; i++ {
		front = front + string([]rune(str)[i])
	}

	return front + str + front
}

