package main

import "fmt"

func main() {
	go func() {
		fmt.Println("go routines")
	}()
	fmt.Println("hello from main")
}
