//https://leetcode.com/problems/container-with-most-water/description/?envType=problem-list-v2&envId=array

package main

import "fmt"

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea2Pointer(arr))
}

// not optimal solution and not ideal for large input values
func maxArea(height []int) int {
	maxInt := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			h := min(height[i], height[j])
			width := j - i
			area := h * width
			if area > maxInt {
				maxInt = area
			}

		}
	}
	return maxInt
}

func maxArea2Pointer(height []int) int {
	left, right := 0, len(height)-1
	max := 0

	for left < right {
		h := min(height[left], height[right])
		width := right - left
		area := h * width
		if area > max {
			max = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}
