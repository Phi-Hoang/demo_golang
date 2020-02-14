package main

import (
	"fmt"
	"math"
)

func main() {
	var f float64
	fmt.Print("Please input a float number: ")
	fmt.Scanf("%f", &f)
	i := int(math.Trunc(f))
	fmt.Printf("Truncated number is: %d", i)
}
