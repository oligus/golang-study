package main

import "fmt"

// We have a loud talking parrot. The "hour" parameter is the current hour time in the range 0..23.
// We are in trouble if the parrot is talking and the hour is before 7 or after 20. Return true if we are in trouble.
//
// parrotTrouble(true, 6) → true
// parrotTrouble(true, 7) → false
// parrotTrouble(false, 6) → false

func main() {
	name := "Parrot trouble"
	fmt.Println(name, "1:", parrotTrouble(true, 6))
	fmt.Println(name, "2:", parrotTrouble(true, 7))
	fmt.Println(name, "3:", parrotTrouble(false, 6))
}

func parrotTrouble(talking bool, hour int) bool {
	if talking && (hour < 7 || hour > 20) {
		return true
	}

	return false
}



