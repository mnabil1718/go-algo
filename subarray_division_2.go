package main

import "fmt"

// source problem: https://www.hackerrank.com/challenges/three-month-preparation-kit-the-birthday-bar/problem
func birthday(s []int32, d int32, m int32) int32 {
	var found int32
	var sum int32

	for i := 0; i < len(s); i++ {

		for j := i; j < int(m)+i; j++ {
			if j < len(s) {
				sum += s[j]
			}
		}

		temp := sum
		sum = 0

		if temp == d {
			found++
		}
	}

	return found
}

func main() {

	var s = []int32{1, 2, 1, 3, 2}

	var found = birthday(s, 3, 2)
	fmt.Println("split ways", found)

}
