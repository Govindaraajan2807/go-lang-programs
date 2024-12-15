package main

import "fmt"

func main() {

	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for value := range intSlice {
		if intSlice[value]%2 == 0 {
			fmt.Println(intSlice[value], " is even")
		} else {
			fmt.Println(intSlice[value], " is odd")
		}
	}
}
