package main

import "fmt"

func main() {
	array := []int{1, 2, 4, 3, 5, 7, 6, 4, 2, 6}
	fmt.Println(findAllDuplicateNumbers(array))
}

func findAllDuplicateNumbers(arr []int) []int {
	var res []int
	for i := 0; i < len(arr); {
		if arr[i] != i+1 {
			correctPosition := arr[i] - 1
			if arr[i] != arr[correctPosition] {
				swap(arr, i, correctPosition)
			} else {
				if !contains(res, arr[i]) {
					res = append(res, arr[i])
				}
				i++
			}
		} else {
			i++
		}
	}
	return res
}

func contains(arr []int, value int) bool {
	for _, val := range arr {
		if val == value {
			return true
		}
	}
	return false
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
