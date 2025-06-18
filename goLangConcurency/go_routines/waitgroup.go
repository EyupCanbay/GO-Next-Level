package main

import (
	"fmt"
	"sync"
)

// go routinelerin çalışırken bir bekleyemeyi olmayınca
// bunu belirtmeyince bu kod çalışmaz canı isterse çalışır
// istemezse çalışmaz
func main() {
	wg := sync.WaitGroup{}
	wg.Add(1) // kaç tane go routine beklemek istediğin

	go func() {
		fmt.Println("hello from goroutine")
		wg.Done()
	}()

	// blocking
	wg.Wait()
	fmt.Println("hello from main")

}
