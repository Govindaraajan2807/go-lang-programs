package main

import "fmt"

func main() {
	// fmt.Print(reverseNumber(123, 0))

	// fmt.Println(palindrome("a"))

	fmt.Println(isPalindrome(123211, 0, 123211))
}

func palindrome(s string) bool {

	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isPalindrome(number, reversed, original int) bool {
	// 1 2 3 2 1
	if number == 0 {
		return reversed == original
	}
	remainder := number % 10
	reversed = reversed*10 + remainder
	return isPalindrome(number/10, reversed, original)

}

func sumOfDigits(n int) int {
	if n == 0 {
		return 0
	}

	return n%10 + sumOfDigits(n/10)
}

func reverseNumber(n, sum int) int {
	//1234
	//4321
	if n == 0 {
		return sum
	}

	rem := n % 10
	sum = sum*10 + rem
	return reverseNumber(n/10, sum)

}
