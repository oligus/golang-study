package main

import (
	"fmt"
	"strings"
)

// Given a string, if the string "del" appears starting at index 1, return a string where that "del" has been deleted.
// Otherwise, return the string unchanged.
//
// delDel("adelbc") → "abc"
// delDel("adelHello") → "aHello"
// delDel("adedbc") → "adedbc"

func main() {
	name := "Del del"
	fmt.Println(name, "1:", delDel("adelbc"))
	fmt.Println(name, "2:", delDel("adelHello"))
	fmt.Println(name, "3:", delDel("adedbc"))
}

func delDel(str string) string {
	if strings.Index(str, "del") == 1 {
		return strings.Replace(str, "del", "", 1)
	}

	return str
}

