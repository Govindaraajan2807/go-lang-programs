package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(arr))
}

func maxSubArray(nums []int) int {
	max := math.MinInt
	currentMax := 0

	for i := 0; i < len(nums); i++ {
		temp := currentMax + nums[i]
		if temp < nums[i] {
			currentMax = nums[i]
		} else {
			currentMax = temp
		}
		if currentMax > max {
			max = currentMax
		}
	}
	return max
}
