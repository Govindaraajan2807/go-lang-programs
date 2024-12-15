package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	fmt.Print(threeSum(arr))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if target == sum {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				// Skip duplicates for the nums[left]
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				// Skip duplicates for nums[right]
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

	return res

	// -4,-1,-1,0,1,2

}
