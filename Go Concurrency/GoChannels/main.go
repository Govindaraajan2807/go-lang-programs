package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	smokeSignal := make(chan bool)
	evilNinjas := "Tommy"
	// smokeSignal <- false // this will create a deadlock and to overcome this we will use buffered channels
	go attack(evilNinjas, smokeSignal)
	fmt.Println(<-smokeSignal)

}

func attack(target string, attacked chan bool) {
	time.Sleep(time.Second)
	fmt.Println("Attacking ninja stars on ", target)
	attacked <- true
}
