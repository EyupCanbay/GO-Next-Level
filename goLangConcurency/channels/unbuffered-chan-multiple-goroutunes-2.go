package main

import (
	"fmt"
	"time"
)

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

	time.Sleep(time.Second)
	fmt.Println("hello chans")

}
