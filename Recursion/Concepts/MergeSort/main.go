package main

import "fmt"

func main() {
	arr := []int{8, 3, 4, 12, 5, 6}
	fmt.Println(mergeSort(arr))
}

func mergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	i, j := 0, 0
	mix := make([]int, 0, len(left)+len(right))

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			mix = append(mix, left[i])
			i++

		} else {
			mix = append(mix, right[j])
			j++
		}
	}

	for i < len(left) {
		mix = append(mix, left[i])
		i++
	}

	for j < len(right) {
		mix = append(mix, right[j])
		j++
	}
	return mix
}
