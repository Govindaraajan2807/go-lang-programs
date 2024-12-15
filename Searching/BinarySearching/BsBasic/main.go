package main

import "fmt"

func main() {
	var arr [6]int = [6]int{1, 2, 3, 4, 5, 6}
	ans := binarySearch(arr, 3)
	fmt.Printf("%d\n", ans)
}

func binarySearch(array [6]int, target int) int { //if size provided, then it will be array
	start := 0            // Initialize start as index 0
	end := len(array) - 1 // Initialize end as the last index

	for start <= end {
		mid := start + (end-start)/2

		if target > array[mid] {
			start = mid + 1
		} else if target < array[mid] {
			end = mid - 1
		} else {
			return array[mid] // Return the found target
		}
	}
	return -1 // Return -1 if target is not found
}
