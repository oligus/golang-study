package main

import (
	"fmt"
)

// Return true if the given string contains between 1 and 3 'e' chars.
//
// stringE("Hello") → true
// stringE("Heelle") → true
// stringE("Heelele") → false

func main() {
	name := "String e"
	fmt.Println(name, "1:", stringE("Hello"))
	fmt.Println(name, "2:", stringE("Heelle"))
	fmt.Println(name, "3:", stringE("Heelele"))
}

func stringE(str string) bool {
	var es int

	for i := 0; i < len(str); i++ {
		letter := string([]rune(str)[i])


		if letter == string('e') {
			es++
		}
	}

	if es >= 1 && es <= 3 {
		return true
	}

	return false
}

