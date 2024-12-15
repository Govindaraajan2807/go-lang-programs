package main

import (
	"fmt"
)

func main() {
	array := []int{2, 4, 1, 6, 8, 5, 10}
	newArray := selectionSort(array)
	fmt.Println(newArray)
}

func selectionSort(array []int) []int {

	for i := 0; i < len(array); i++ {
		lastElement := len(array) - 1 - i
		maxIndex := getMaxIndex(array, 0, lastElement)
		swap(array, maxIndex, lastElement)
	}
	return array
}

func getMaxIndex(array []int, start, end int) int {
	max := start

	for i := start; i <= end; i++ {
		if array[i] > array[max] {
			max = i
		}
	}
	return max
}

func swap(array []int, i, j int) {
	temp := array[i]
	array[i] = array[j]
	array[j] = temp
}
