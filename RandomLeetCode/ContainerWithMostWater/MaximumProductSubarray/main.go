package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(arr))
}

func maxProduct(nums []int) int {

	prefix, suffix, result := 1, 1, math.MinInt

	for i := 0; i < len(nums); i++ {
		prefix *= nums[i]
		suffix *= nums[len(nums)-1-i]

		result = int(math.Max(float64(result), math.Max(float64(prefix), float64(suffix))))

		if prefix == 0 {
			prefix = 1
		}
		if suffix == 0 {
			suffix = 1
		}

	}
	return result

}
