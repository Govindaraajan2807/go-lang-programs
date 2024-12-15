// https://leetcode.com/problems/4sum/description/?envType=problem-list-v2&envId=array

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-3, -1, 0, 2, 4, 5}
	fmt.Print(fourSum(arr, 0))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	n := len(nums)

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, n-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if target == sum {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return res
}
