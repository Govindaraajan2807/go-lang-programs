package main

import (
	"fmt"
	"sync"
)

func main() {

	var beeper sync.WaitGroup
	evilNinjas := []string{"Tommy", "Bobby", "Andy"}
	beeper.Add(len(evilNinjas))
	for _, evilNinja := range evilNinjas {
		go attack(evilNinja, &beeper)
	}
	beeper.Wait()
	fmt.Println("Mission accomplished")

}

func attack(evilNinja string, beeper *sync.WaitGroup) {
	fmt.Println("Attacking ninja star on: ", evilNinja)
	beeper.Done()
}
