// https://leetcode.com/problems/missing-number/submissions/1451647118/
package main

import "fmt"

func main() {
	arr := []int{4, 0, 1, 2, 7, 3, 8, 5}
	fmt.Println("The missing number is: ", missingNumber(arr))
}

func missingNumber(arr []int) int {

	n := len(arr)

	for i := 0; i < n; {
		correctPosition := arr[i] //correctPos = 4
		if arr[i] < n && arr[correctPosition] != arr[i] {
			swap(arr, i, correctPosition)
		} else {
			i++
		}
	}

	for i := 0; i < n; i++ {
		if arr[i] != i {
			return i
		}
	}
	return -1

}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
