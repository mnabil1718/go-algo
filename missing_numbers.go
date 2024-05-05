package main

import (
	"fmt"
	"sort"
)

// source: https://www.hackerrank.com/challenges/three-month-preparation-kit-missing-numbers/problem
func missingNumbers(arr []int32, brr []int32) []int32 {

	var brr_map map[int32]int32 = make(map[int32]int32)

	// making map from brr
	for i := 0; i < len(brr); i++ {
		if value, ok := brr_map[brr[i]]; ok {
			brr_map[brr[i]] = value + 1
		} else {
			brr_map[brr[i]] = 1
		}
	}

	// checking arr values in brr_map
	for i := 0; i < len(arr); i++ {
		if value, ok := brr_map[arr[i]]; ok {
			brr_map[arr[i]] = value - 1
		}
	}

	var resultArray []int32
	for key, value := range brr_map {
		if value > 0 {
			resultArray = append(resultArray, key)
		}
	}

	sort.Slice(resultArray, func(i, j int) bool { return resultArray[i] < resultArray[j] })

	return resultArray

}

func main() {
	arr := []int32{203, 204, 205, 206, 207, 208, 203, 204, 205, 206}
	brr := []int32{203, 204, 204, 205, 206, 207, 205, 208, 203, 206, 205, 206, 204}
	resultArray := missingNumbers(arr, brr)
	fmt.Println(resultArray)
}
