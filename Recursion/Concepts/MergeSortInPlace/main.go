package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1}
	mergeSortInPlace(arr, 0, len(arr))
	fmt.Println(arr)
}

func mergeSortInPlace(arr []int, s, e int) {
	if e-s == 1 {
		return
	}

	m := s + (e-s)/2

	mergeSortInPlace(arr, s, m)
	mergeSortInPlace(arr, m, e)

	merge(arr, s, m, e)
}

func merge(arr []int, s, m, e int) []int {
	i, j := s, m
	mix := make([]int, 0, e-s)
	for i < j && j < e {
		if arr[i] < arr[j] {
			mix = append(mix, arr[i])
			i++
		} else {
			mix = append(mix, arr[j])
			j++
		}
	}
	for i < m {
		mix = append(mix, arr[i])
		i++
	}

	for j < e {
		mix = append(mix, arr[j])
		j++
	}

	for k := 0; k < len(mix); k++ {
		arr[s+k] = mix[k]
	}
	return arr
}
