package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//raceExample()
	//raceExampleFixed()
	raceExampleFixedWithAtomic()
}

// faklı go routinelerin aynı memorye erişeye çalışmasıdır
func raceExample() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	// shared value
	val := 0

	go func() {
		for i := 0; i < 100000; i++ {
			val++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100000; i++ {
			val++
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(val)

}

func raceExampleFixed() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	// eş zamanlı olarak sadece bir routinenin erişebilceğini
	//garanti altına alır
	lock := sync.Mutex{}

	val := 0

	go func() {
		for i := 0; i < 100000; i++ {
			lock.Lock()
			val++
			lock.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100000; i++ {
			lock.Lock()
			val++
			lock.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(val)
}

func raceExampleFixedWithAtomic() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	var val int32 = 0

	go func() {
		for i := 0; i < 100000; i++ {
			atomic.AddInt32(&val, 1)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100000; i++ {
			atomic.AddInt32(&val, 1)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(val)
}
