//https://leetcode.com/problems/combination-sum/?envType=problem-list-v2&envId=array&difficulty=MEDIUM

package main

import "fmt"

func main() {
	arr := []int{2, 3, 5}
	target := 8
	fmt.Println(combinationSum(arr, target))
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	n := len(candidates)

	for i := 0; i < n; i++ {
		// 1. If element exists in array
		if target == candidates[i] {
			res = append(res, []int{candidates[i]})
			continue
		}

		// 2. Check if element is multiple of target
		if target%candidates[i] == 0 {
			counter := 1
			candidate := candidates[i]
			value := candidates[i] //2

			for value < target { //2<8
				value = candidates[i] * counter
				fmt.Println("Value is : ", value)
				counter++
				fmt.Println("Counter value inside loop : ", counter)
			}
			if value == target {
				fmt.Println("Counter value is : ", counter)
				for i := 1; i < counter; i++ {
					res = append(res, []int{candidate})
				}
			}
			continue
		}

		// 3. take a element and multiply it until <  target and find a sum
		element := candidates[i]
		actualElement := element
		counter := 1

		for element <= target-element {
			counter++
			element += element
		}
		find := target - element
		for j := 0; j < n; j++ {
			if candidates[j] == find {
				res = append(res, []int{actualElement, candidates[j]})
			}
		}
	}

	return res
}
