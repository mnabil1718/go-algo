package main

import (
	"fmt"
	"math"
	"sort"
)

// source: https://www.hackerrank.com/challenges/three-month-preparation-kit-closest-numbers/problem
func closestNumbers(arr []int32) []int32 {
	// Write your code here
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	minAbsDiff := int(math.Abs(float64(math.Abs(float64(arr[1])) - math.Abs(float64(arr[0])))))
	var storage []int32 = []int32{arr[0], arr[1]}

	if len(arr) == 2 {
		return storage
	}

	for i := 1; i < len(arr)-1; i++ {
		currentDiff := int(math.Abs(float64(math.Abs(float64(arr[i+1])) - math.Abs(float64(arr[i])))))

		if currentDiff < minAbsDiff {
			minAbsDiff = currentDiff
			storage = []int32{}
			storage = append(storage, arr[i], arr[i+1])

		} else if currentDiff == minAbsDiff {
			storage = append(storage, arr[i], arr[i+1])

		}
	}

	return storage

}

func main() {

	// a := []int32{5, 2, 3, 4, 1}
	b := []int32{-20, -3916237, -357920, -3620601, 7374819, -7330761, 30, 6246457, -6461594, 266854}

	arr := closestNumbers(b)
	fmt.Println(arr)

}
