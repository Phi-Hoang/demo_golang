package main

import (
	"fmt"
	"time"
)

func main() {
	go Task1()
	go Task2()
	Task3()
}

func Task1() {
	time.Sleep(2 * time.Second)
	fmt.Println("task 1 done")
}

func Task2() {
	time.Sleep(2 * time.Second)
	fmt.Println("task 2 done")
}

func Task3() {
	time.Sleep(2*time.Second + 100*time.Millisecond)
	fmt.Println("task 3 done")
}
