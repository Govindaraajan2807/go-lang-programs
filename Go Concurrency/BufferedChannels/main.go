package main

import "fmt"

func main() {
	channel := make(chan string, 2) // mentioning the size ",2" will make it as a buffered channel
	channel <- "First message"
	channel <- "Second message"
	// channel <- "Third message" // Having it here with size 2 will throw deadlock since 2 messages are not spit out from the channel
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	channel <- "Third message" // since the 1st 2 messages were split out, channel has empty space and 3rd msg is consumed
	fmt.Println(<-channel)
}
