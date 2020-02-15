package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	// print hello 1, hello 2, hello 3,... in every 800ms in go routine
	//** remember keep main routine alive (active/sleep -> not finished)
	forever := make(chan os.Signal)
	signal.Notify(forever, os.Interrupt)
	go func() {
		i := 0
		for true {
			i++
			fmt.Println("hello", i)
			time.Sleep(time.Millisecond * 800)
		}
	}()
	<-forever
	fmt.Println("Program end")
}
