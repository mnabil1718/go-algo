package main

import "fmt"

// source problem see: https://www.hackerrank.com/challenges/flipping-the-matrix/problem

func flipMatrix(matrix [][]int) int {
	var n int = len(matrix) / 2

	var sum int = 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum += max(matrix[i][j], matrix[i][2*n-1-j], matrix[2*n-1-i][2*n-1-j], matrix[2*n-1-i][j])
		}
	}

	return sum
}

func main() {

	var arr = [][]int{
		{112, 42, 83, 119},
		{56, 125, 56, 49},
		{15, 78, 101, 43},
		{62, 98, 114, 108},
	}

	var sum = flipMatrix(arr)
	fmt.Println("sum", sum)

}
