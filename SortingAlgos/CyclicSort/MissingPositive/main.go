package main

import "fmt"

func main() {
	array := []int{1, 8, 7, 9, -1, 0, 2}
	fmt.Println(missingPositive(array))
}

func missingPositive(arr []int) int {
	n := len(arr)

	for i := 0; i < n; {
		correctPosition := arr[i] - 1

		if arr[i] > 0 && arr[i] <= n && arr[i] != arr[correctPosition] {
			swap(arr, i, correctPosition)
		} else {
			i++
		}

	}

	for i := 0; i < n; i++ {
		if arr[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
