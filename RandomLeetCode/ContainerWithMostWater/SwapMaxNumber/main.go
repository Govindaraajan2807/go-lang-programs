package main

import "fmt"

func main() {

	fmt.Println(maximumSwap(2736))

}

func maximumSwap(num int) int {
	//convert to integer array
	arr := []int{}
	duplicate := num

	for duplicate > 0 {
		digit := duplicate % 10
		arr = append([]int{digit}, arr...)
		duplicate /= 10
	}

	imap := make(map[int]int)
	for i, digit := range arr {
		imap[digit] = i
	}

	for i, digit := range arr {
		for d := 9; d > digit; d-- {
			if pos, exists := imap[d]; exists && pos > i {
				arr[i], arr[pos] = arr[pos], arr[i]
				return arrayToInt(arr)
			}
		}
	}

	return num

}

func arrayToInt(arr []int) int {
	res := 0

	for _, digit := range arr {
		res = res*10 + digit
	}
	return res
}
