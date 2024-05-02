package main

import (
	"fmt"
	"strings"
)

// source: https://www.hackerrank.com/challenges/anagram/problem
func anagram(s string) int32 {

	var count int32 = 0

	// Write your code here
	if len(s)%2 != 0 {
		return -1
	}

	left, right := s[:len(s)/2], s[len(s)/2:]

	for i := 0; i < len(left); i++ {

		if !strings.Contains(right, string(left[i])) {
			count++
		} else {
			right = strings.Replace(right, string(left[i]), "", 1)
		}
	}

	return count

}

func main() {
	sample := "aaabbb" // 3
	// ab = 1
	// abc = -1
	// mnop = 2
	// xyyx = 0
	// xaxbbbxx = 1
	// fdhlvosf pafhalll = 5
	fmt.Println(anagram(sample))
}
