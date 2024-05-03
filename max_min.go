package main

import (
	"fmt"
	"math"
	"sort"
)

// source: https://www.hackerrank.com/challenges/three-month-preparation-kit-angry-children/problem
func maxMin(k int32, arr []int32) int32 {

	var min int32 = math.MaxInt32
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for i := 0; i <= len(arr)-int(k); i++ {
		diff := arr[i+int(k)-1] - arr[i]
		if min > diff {
			min = diff
		}
	}

	return min
}

func main() {
	sample := []int32{1, 4, 7, 2}
	fmt.Println(maxMin(2, sample))
}
