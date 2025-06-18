package main

import "fmt"

func main() {
	chan1 := make(chan int, 1)
	chan1 <- 1

	select {
	case val := <-chan1:
		fmt.Println(val)
	}
	
}
