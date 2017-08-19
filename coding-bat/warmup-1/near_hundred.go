package main

import (
	"fmt"
	"math"
)

// Given an int n, return true if it is within 10 of 100 or 200.
// Note: Math.abs(num) computes the absolute value of a number.
//
// nearHundred(93) → true
// nearHundred(90) → true
// nearHundred(89) → false

func main() {
	name := "Near hundred"
	fmt.Println(name, "1:", nearHundred(93))
	fmt.Println(name, "2:", nearHundred(90))
	fmt.Println(name, "3:", nearHundred(89))
}

func nearHundred(n int) bool {
	return math.Abs(100 - float64(n)) <= 10 || math.Abs(200 - float64(n)) <= 10
}


