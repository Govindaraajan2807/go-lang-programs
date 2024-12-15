package main

import (
	"fmt"
	"math/rand"
)

func main() {
	channel := make(chan string)
	go throwingStars(channel)
	for {
		message, open := <-channel
		if !open {
			break
		}
		fmt.Println(message)
	}
}

func throwingStars(channel chan string) {

	chances := 3
	for i := 0; i < chances; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("You have scored: ", score)
	}
	close(channel) // closing the channel after feeding in value to avoid deadlock in main go routine
}
