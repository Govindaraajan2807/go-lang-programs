package main

import (
	"fmt"
	"time"
)

func main() {
	ninja1, ninja2 := make(chan string), make(chan string)
	// ninja1, ninja2 := make(chan interface), make(chan string)

	go captainElect(ninja1, "Ninja1")
	go captainElect(ninja2, "Ninja2")

	select {
	case message := <-ninja1:
		fmt.Println(message)
	case message := <-ninja2:
		fmt.Println(message)
	default:
		fmt.Println("Neither..")
	}

	roughlyFair()

	// fmt.Println(<-ninja1)
	// fmt.Println(<-ninja2)
}

func roughlyFair() {
	ninja1 := make(chan interface{})
	close(ninja1)
	ninja2 := make(chan interface{})
	close(ninja2)

	var ninjaCount1, ninjaCount2 int

	for i := 0; i < 1000; i++ {
		select {
		case <-ninja1:
			ninjaCount1++
		case <-ninja2:
			ninjaCount2++
		}
	}

	fmt.Printf("ninjaCount1 : %d, ninjaCount2 : %d ", ninjaCount1, ninjaCount2)
}

func captainElect(ninja chan string, message string) {
	time.Sleep(1 * time.Second)
	ninja <- message
}
