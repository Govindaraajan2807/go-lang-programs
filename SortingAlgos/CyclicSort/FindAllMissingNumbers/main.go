// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/submissions/1451671136/
package main

import "fmt"

func main() {
	arr := []int{8, 1, 2, 7, 3, 9, 5, 6, 0}
	fmt.Println("Missing numbers:", findMissingNumbers(arr))
}

func findMissingNumbers(arr []int) []int {
	//arr := []int{8, 1, 2, 7, 3, 9, 5, 6, 0}
	n := len(arr)
	fmt.Println("Length of array is : ", n)
	for i := 0; i < n; {
		correctPosition := arr[i]
		if arr[i] < n && arr[i] >= 0 && arr[correctPosition] != arr[i] {
			swap(arr, i, correctPosition)
		} else {
			i++
		}
	}

	fmt.Println("Sorted array is : ", arr)

	var resArray []int

	// 1, 2, 3, 4, 5, 7, 8
	for i := 0; i < n; i++ {
		if arr[i] != i {
			resArray = append(resArray, i)
		}
	}
	return resArray
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
