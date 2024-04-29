package main

import (
	"fmt"
	"strconv"
)

// source: https://www.hackerrank.com/challenges/three-month-preparation-kit-separate-the-numbers/problem
func separateNumbers(s string) {
	// verbose version to see how it works
	for i := 1; i <= len(s)/2; i++ {
		start := s[:i]
		fmt.Println("start", start)
		num, _ := strconv.Atoi(start)
		fmt.Println("num", num)
		var built string = ""

		for len([]rune(built)) < len([]rune(s)) {
			num++
			fmt.Println("num increased", num)
			built += strconv.Itoa(num)
			fmt.Println("built", built)

			if built == s[i:] {
				fmt.Println("YES", start)
				return
			}
		}
	}

	fmt.Println("NO")

}

func main() {

	numbers := "91011"

	separateNumbers(numbers)

}
