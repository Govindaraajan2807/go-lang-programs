package main

import (
	"fmt"
	"time"
)

func main() {
	evilNinjas := []string{"Tommy", "Bobby", "Johnny", "Andy"}
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	for _, evilNinja := range evilNinjas {
		go attack(evilNinja)
	}

	time.Sleep(time.Second * 1)
}

func attack(target string) {
	fmt.Println("Throwing ninja stars at ", target)
	time.Sleep(time.Second)
}
