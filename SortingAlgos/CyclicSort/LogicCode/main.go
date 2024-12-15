package main

import "fmt"

func main() {
	arr := []int{4, 1, 6, 2, 7, 3, 8, 5}
	fmt.Println(cyclicSort(arr))
}

func cyclicSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		if arr[i] >= 1 && arr[i] <= len(arr) && arr[i] != arr[arr[i]-1] {
			swap(arr, i, arr[i]-1)
		} else {
			i++
		}
	}
	return arr
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
