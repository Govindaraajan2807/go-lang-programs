package main

import "fmt"

func main() {
	array := []int{1, 2, 4, 3, 5, 7, 6, 4}
	fmt.Print(findDuplicateNumber(array))
}

func findDuplicateNumber(arr []int) int {
	n := len(arr)

	//{1, 2, 4, 3, 5, 7, 6,4}
	for i := 0; i < n; {
		if arr[i] != i+1 {
			correctPosition := arr[i] - 1
			if arr[i] != arr[correctPosition] {
				swap(arr, i, correctPosition)
			} else {
				return arr[i]
			}
		} else {
			i++
		}
	}
	/*
		for i := 0; i < n; {
			correctPosition := arr[i] - 1
			if arr[i] != arr[correctPosition] {
				swap(arr, i, correctPosition)
			} else {
				i++
			}
		}

		for i := 0; i < n; i++ {
			if arr[i] != i+1 {
				return arr[i]
			}
		}
	*/

	return -1
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
