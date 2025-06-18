package main

import "fmt"

func main() {

	//channel initialization
	unbufferedChan := make(chan int)

	//reader goroutine
	go func(unbufChan chan int) {
		value := <-unbufferedChan
		fmt.Println("value from unbuffered chan:", value)
	}(unbufferedChan)

	//writer goroutine
	go func(unbufChan chan int) {
		unbufferedChan <- 1
	}(unbufferedChan)

	fmt.Println("hello chans")

	/*
			output is non-deterministic. scheduler probably will not have
		time to schedule goroutines. So we will not see channel value in the output

	*/
}
