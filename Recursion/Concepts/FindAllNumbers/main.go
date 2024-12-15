package main

import "fmt"

func main() {
	// arr := []int{1, 2, 3, 4, 4, 5, 6}
	// fmt.Println(findAllIndex(arr, 0, 4, []int{}))
	// fmt.Println(findAllIndex2(arr, 0, 4))

	arr := []int{1, 2, 3, 2, 4, 2, 5}
	target := 2

	result := findAll(arr, target, 0)
	fmt.Println("Indices of target:", result)
}

func findAllIndex(arr []int, index int, target int, res []int) []int {
	if index == len(arr) {
		return res
	}

	if target == arr[index] {
		res = append(res, index)
	}
	return findAllIndex(arr, index+1, target, res)
}

func findAll(arr []int, target, index int) []int {
	if index == len(arr) {
		return []int{}
	}

	// Recursive call to find indices in the rest of the array.
	addBelowListValues := findAll(arr, target, index+1)

	// If the current element matches the target, add the index to the result.
	if arr[index] == target {
		return append([]int{index}, addBelowListValues...)
	}

	// Otherwise, return the result from the recursive call.
	return addBelowListValues
}
