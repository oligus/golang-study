package main

import (
	"fmt"
	"strings"
)

// Return true if the given string begins with "mix", except the 'm' can be anything, so "pix", "9ix" .. all count.
//
// mixStart("mix snacks") → true
// mixStart("pix snacks") → true
// mixStart("piz snacks") → false
// mixStart("ix snacks") → true

func main() {
	name := "Mix start"
	fmt.Println(name, "1:", mixStart("mix snacks"))
	fmt.Println(name, "2:", mixStart("pix snacks"))
	fmt.Println(name, "3:", mixStart("piz snacks"))
	fmt.Println(name, "4:", mixStart("ix snacks"))
}

func mixStart(str string) bool {
	position := strings.Index(str, "ix")

	if position == 1 || position == 0 {
		return true
	}

	return false
}

