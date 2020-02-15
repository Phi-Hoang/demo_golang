package main

func main() {
	ch := make(chan int)

	// read from channel cause routine sleep
	<-ch
}
