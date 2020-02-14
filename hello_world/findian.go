package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Print("Please enter a string: ")
	fmt.Scanf("%s", &s)
	if strings.Contains(s, "i") && strings.Contains(s, "a") && strings.Contains(s, "n") {
		fmt.Print("Found!")
	} else {
		fmt.Print("Not Found!")
	}
}
