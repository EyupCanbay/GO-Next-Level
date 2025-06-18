package main

import "fmt"

func main() {
	// channel initialization
	unbufferedChan := make(chan int)

	// channel declaration
	var unbuffered2 chan int

	fmt.Println("unbufferedChan:", unbufferedChan)
	fmt.Println("unbuffered2:", unbuffered2)

	//only read from channel
	go func(unbufChan <-chan int) {
		// blocks until data arrives
		value := <-unbufChan
		fmt.Println("value =", value)
	}(unbufferedChan)

	unbufferedChan <- 1
}
