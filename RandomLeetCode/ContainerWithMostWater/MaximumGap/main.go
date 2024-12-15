package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 10000000}
	fmt.Println(maximumGap(arr))
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	max := 0
	for i := 0; i < len(nums)-1; i++ {
		j := i + 1
		gap := nums[j] - nums[i]
		if gap > max {
			max = gap
		}
	}
	return max
}
