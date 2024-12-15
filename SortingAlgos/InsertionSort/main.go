package main

import "fmt"

func insertionSort(array []int) []int {
	// Traverse the array from the second element to the end
	for i := 1; i < len(array); i++ {
		key := array[i] // Element to be inserted
		j := i - 1      // The last index of the sorted part

		// Shift elements of array[0..i-1] that are greater than key to one position ahead
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j--
		}

		// Insert the key at the correct position
		array[j+1] = key
	}
	return array
}

func main() {
	array := []int{5, 2, 9, 1, 5, 6}
	sortedArray := insertionSort(array)
	fmt.Println(sortedArray) // Output: [1 2 5 5 6 9]
}
