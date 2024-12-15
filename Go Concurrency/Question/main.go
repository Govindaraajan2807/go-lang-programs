package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {

	ch := make(chan bool, 2)
	ch <- true
	ch <- true
	fmt.Print(<-ch)
	close(ch)
	ch <- false

}

func locker() {
	N := 100
	m := make(map[int]bool)
	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			r := rand.Intn(1000)
			mu.Lock()
			m[r] = true
			mu.Unlock()
		}()
	}
	wg.Wait()

	for _, val := range m {
		fmt.Println(val)
	}
}
