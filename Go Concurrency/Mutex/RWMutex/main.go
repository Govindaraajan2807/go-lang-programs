package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count  int
	lock   sync.Mutex
	rwLock sync.RWMutex
)

func main() {
	basics()
}

func readAndWrite() {
	go read()
	go read()
	go read()
	go read()
	go read()
	go write()

	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}

func read() {
	rwLock.RLock()
	defer rwLock.RUnlock()

	fmt.Println("Read locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Read unlocking")

}

func write() {
	rwLock.Lock()
	defer rwLock.Unlock()

	fmt.Println("Write locking")
	time.Sleep(1 * time.Second)
	fmt.Println("Write unlocking")
}

func basics() {
	counter := 1000
	for i := 0; i < counter; i++ {
		go increment()
	}
	// time.Sleep(1 * time.Second)
	fmt.Println("Value of count is: ", count)
}

func increment() {
	lock.Lock()
	count++
	lock.Unlock()
}
