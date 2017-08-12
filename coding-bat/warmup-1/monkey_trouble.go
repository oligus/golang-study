package main

import "fmt"

// We have two monkeys, a and b, and the parameters aSmile and bSmile indicate if each is smiling.
// We are in trouble if they are both smiling or if neither of them is smiling. Return true if we are in trouble.
//
// monkeyTrouble(true, true) → true
// monkeyTrouble(false, false) → true
// monkeyTrouble(true, false) → false

func main() {
	fmt.Println("Monkey Trouble 1:", monkeyTrouble(true, true))
	fmt.Println("Monkey Trouble 2:", monkeyTrouble(false, false))
	fmt.Println("Monkey Trouble 3:", monkeyTrouble(true, false))
}

func monkeyTrouble(aSmile, bSmile bool) bool {
	return aSmile == bSmile
}


