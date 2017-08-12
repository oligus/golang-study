package main

import "fmt"

// The parameter weekday is true if it is a weekday, and the parameter vacation is true if we are on vacation.
// We sleep in if it is not a weekday or we're on vacation. Return true if we sleep in.
//
// sleepIn(false, false) → true
// sleepIn(true, false) → false
// sleepIn(false, true) → true

func main() {
	fmt.Println("Sleep in 1:", sleepIn(false, false))
	fmt.Println("Sleep in 2:", sleepIn(true, false))
	fmt.Println("Sleep in 3:", sleepIn(false, true))
}

func sleepIn(weekday, vacation bool) bool {
	return !weekday || vacation
}

