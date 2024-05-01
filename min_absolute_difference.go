package main

import (
	"fmt"
	"math"
	"sort"
)

// source: https://www.hackerrank.com/challenges/three-month-preparation-kit-minimum-absolute-difference-in-an-array/problem
func minimumAbsoluteDifference(arr []int32) int32 {

	// if you do benchmarking this operation
	// actually took longer than without sorting
	// on len(arr) = 10 test case, but once you
	// go bigger like len(arr) = 100000, then
	// sorting overhead is worth it
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	var min int32 = math.MaxInt32

	// Write your code here
	for i := 0; i < len(arr)-1; i++ {

		diff := int32(math.Abs(float64(arr[i] - arr[i+1])))

		if diff < min {
			min = diff
		}
	}

	return min
}

func main() {
	arr := []int32{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75}
	fmt.Println(minimumAbsoluteDifference(arr))
}
