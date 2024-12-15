package main

import "fmt"

func main() {
	array := [4][4]int{
		{10, 20, 30, 40},
		{29, 25, 35, 45},
		{28, 15, 37, 49},
		{33, 34, 38, 50},
	}

	fmt.Println(searchIn2DArray(array, 25))

}

func searchIn2DArray(array [4][4]int, target int) []int {
	r := 0
	c := len(array[0]) - 1

	for r < len(array) && c >= 0 {
		if target == array[r][c] {
			return []int{r, c}
		} else if array[r][c] > target {
			c--
		} else {
			r++
		}
	}
	return []int{-1, -1}
}
