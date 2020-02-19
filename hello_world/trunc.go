package main

import (
	"fmt"
	"math"
)

/*
 * Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered.
 * Truncation is the process of removing the digits to the right of the decimal place.
 */
func main() {
	var f float64
	fmt.Print("Please input a float number: ")
	fmt.Scanf("%f", &f)
	i := int(math.Trunc(f))
	fmt.Printf("Truncated number is: %d", i)
}
