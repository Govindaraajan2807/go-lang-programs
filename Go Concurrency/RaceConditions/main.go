package main

import (
	"fmt"
	"sync"
)

var counter int
var muLock sync.Mutex

func main() {
	var wg sync.WaitGroup

	// Spawn 10 goroutines to increment the counter
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Final Counter:", counter)
}

func increment(wg *sync.WaitGroup) {

	for i := 0; i < 1000; i++ {
		muLock.Lock() // without mutex lock, it will lead to race conditions
		counter++
		muLock.Unlock()
	}
	defer wg.Done()
}
