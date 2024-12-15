package main

import "fmt"

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Returning from  a panic")
		}
	}()

	fmt.Println("About to panic")
	panic("Returning from panic")
	// fmt.Println("Returning from  main") //this line will not be executed
}
