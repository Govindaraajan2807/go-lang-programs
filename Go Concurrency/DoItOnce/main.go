package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var missionCompleted bool

var once sync.Once

func main() {

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			if foundTreasure() {
				once.Do(markAsCompleted)
				// markAsCompleted()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	checkMissionStatus()
}

func checkMissionStatus() {
	if missionCompleted {
		fmt.Println("Mision is completed..")
	} else {
		fmt.Println("Mission was a failure")
	}
}

func foundTreasure() bool {
	rand.Seed(time.Now().Unix())
	return 0 == rand.Intn(10)
}

func markAsCompleted() {
	missionCompleted = true
	fmt.Println("Marking mission as completed")
}
