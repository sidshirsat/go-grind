// https://leetcode.com/problems/valid-palindrome/

package main

import "fmt"

func validPalindrome(str string) bool {

	left := 0
	right := len(str) - 1

	for left <= right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(validPalindrome("abxsba"))
}
