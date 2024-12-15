package main

import "fmt"

func main() {
	arr := []int{1}
	target := 1

	var first, second int

	first = firstLast(arr, target, true)
	second = firstLast(arr, target, false)
	fmt.Println(first, second)

	// value := binarySearch(arr, target)
	// fmt.Println(value)

	// if value > 0 {
	// 	first = firstAndLast(arr, 0, value, target)
	// 	second = firstAndLast(arr, value, len(arr)-1, target)
	// } else if value == 0 {
	// 	first, second = 0, 0
	// } else {
	// 	first, second = -1, -1
	// }
	// fmt.Println(first, second)
}

func firstLast(arr []int, target int, isFirst bool) int {
	start := 0
	end := len(arr) - 1
	result := -1

	for start <= end {
		mid := start + (end-start)/2
		if target > arr[mid] {
			start = mid + 1
		} else if target < arr[mid] {
			end = mid - 1
		} else {
			result = mid
			if isFirst {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	return result
}
