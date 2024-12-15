package main

import "fmt"

func main() {
	array := []int{2, 3, 1, 2, 5, 4}
	fmt.Println(cyclicSort(array))
}

func cyclicSort(arr []int) []int {

	n := len(arr)

	for i := 0; i < n; {
		correctPosition := arr[i] - 1
		if arr[i] != arr[correctPosition] {
			swap(arr, i, correctPosition)
		} else {
			i++
		}
	}

	fmt.Println("Sorted array is : ", arr)

	var res []int

	for i := 0; i < n; i++ {
		if arr[i] != i+1 {
			res = append(res, arr[i], i+1)
		}
	}
	return res
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
