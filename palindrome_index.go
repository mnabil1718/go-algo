package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

// source: https://www.hackerrank.com/challenges/palindrome-index/problem
func palindromeIndex(s string) int32 {
	var left, right int = 0, len(s) - 1

	for left < right {
		if string(s[left]) != string(s[right]) {
			if isPalindrome(s[left+1 : right+1]) {
				return int32(left)
			} else if isPalindrome(s[left:right]) {
				return int32(right)
			} else {
				return -1
			}
		}

		right--
		left++
	}

	return -1

}

func main() {
	test_cases := []string{
		"bcbc", // 0 or 3
		"aaab", // 3
		"baa",  // 0
		"aaa",  // -1 (already palindrome)
	}

	for _, test_case := range test_cases {
		fmt.Println(palindromeIndex(test_case))
	}
}
