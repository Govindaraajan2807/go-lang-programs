package main

import "fmt"

func main() {

	cars := map[string]string{
		"bmw":  "M3",
		"audi": "a6",
		"benz": "amg",
	}

	cars = modifyMap(cars)
	printMap(cars)
}

func printMap(m map[string]string) {
	for make, model := range m {
		fmt.Println(make, model)
	}
}

func modifyMap(m map[string]string) map[string]string {
	tempMap := make(map[string]string)

	for make, model := range m {
		if make == "bmw" {
			model = "M4"
		}
		tempMap[make] = model
	}
	return tempMap
}
