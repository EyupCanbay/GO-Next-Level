package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "", "specify your name")
	flag.Parse()

	if len(name) == 0 {
		fmt.Println("you must specify a name")
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Printf("Hello %s/n", name)
}
