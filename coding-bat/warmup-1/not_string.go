package main

import (
	"fmt"
	"strings"
)

// Given a string, return a new string where "not " has been added to the front.
// However, if the string already begins with "not", return the string unchanged. Note: use .equals() to compare 2 strings.
//
// notString("candy") → "not candy"
// notString("x") → "not x"
// notString("not bad") → "not bad"

func main() {
	name := "Not string"
	fmt.Println(name, "1:", notString("candy"))
	fmt.Println(name, "2:", notString("x"))
	fmt.Println(name, "3:", notString("not bad"))
}

func notString(str string) string {

	if strings.HasPrefix(str, "not") {
		return str
	}

	return "not " + str
}


