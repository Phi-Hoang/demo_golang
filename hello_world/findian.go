package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	fmt.Print("Please enter a string: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s = scanner.Text()
	}
	if strings.HasPrefix(strings.ToLower(s), "i") && strings.Contains(strings.ToLower(s), "a") && strings.HasSuffix(strings.ToLower(s), "n") {
		fmt.Print("Found!")
	} else {
		fmt.Print("Not Found!")
	}
}
