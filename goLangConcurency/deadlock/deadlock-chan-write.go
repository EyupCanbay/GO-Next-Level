package main

import "fmt"

func main() {
	// unbuffered channel
	bufChan := make(chan int)

	// write are blocking operations
	bufChan <- 1

	fmt.Println("hello")
}
