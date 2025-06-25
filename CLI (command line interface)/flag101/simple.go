package main

import (
	"flag"
	"fmt"
)

// <program> <flags> <argument>
// ./simple -flavor	"abc" -quantity 3 -crema

// mini order cli
func main() {
	wordPtr := flag.String("flavor", "vanilla", "select shot flavor")
	numPtr := flag.Int("quantity", 10, "select of short")
	boolPtr := flag.Bool("crema", false, "decide if you want crema")

	var order string
	flag.StringVar(&order, "order", "complete", "status of order")
	flag.Parse() // if you forget to parse, it did not accept your enter the comment line argument

	fmt.Println("wordPtr =", *wordPtr)
	fmt.Println("numPtr =", *numPtr)
	fmt.Println("boolPtr =", *boolPtr)
	fmt.Println("order", order)
	fmt.Println("tail", flag.Args())

}
