package main

import (
	"fmt"
	"sync"
)

type RaceTest struct {
	Val int
}

func main() {
	raceTest := &RaceTest{}

	wg := sync.WaitGroup{}
	wg.Add(10000)

	lock := sync.Mutex{}
	for i := 0; i < 10000; i++ {

		go increment(raceTest, &wg, &lock)
	}
	wg.Wait()
 
	fmt.Println(raceTest)
}
func increment(rt *RaceTest, wg *sync.WaitGroup, l *sync.Mutex) {
	l.Lock()
	rt.Val++
	l.Unlock()
	wg.Done()
}
