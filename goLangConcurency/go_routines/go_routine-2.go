package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		go func(value int) { // value olarak aluyoruz
			fmt.Println("hello from go routine", value)
		}(i) // i parametresini gopy referans olarak göneriyoruz
	}

	time.Sleep(3 * time.Second)
}
