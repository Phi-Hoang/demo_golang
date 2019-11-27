package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// sync.WaitGroup: wg.Add() wg.Done() wg.Wait()
	var wg sync.WaitGroup
	wg.Add(2)
	go Task1(&wg)
	go Task2(&wg)
	Task3()
	wg.Wait()
}

func Task1(wg *sync.WaitGroup) {
	defer wg.Done() // lastly run
	time.Sleep(2 * time.Second)
	fmt.Println("task 1 done")
}

func Task2(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("task 2 done")
	wg.Done()
}

func Task3() {
	fmt.Println("task 3 done")
}
